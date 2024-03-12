package templating

import (
	"github.com/Drafteame/taskrun/internal/models"
)

func (jt *JobTemplate) Render() (*models.JobConfig, error) {
	if err := jt.renderTemplateData(); err != nil {
		return nil, err
	}

	if errFrom := jt.jobModel.FromYAML([]byte(jt.template)); errFrom != nil {
		return nil, errFrom
	}

	if errRemote := jt.renderRemote(); errRemote != nil {
		return nil, errRemote
	}

	if err := jt.renderDynamic(); err != nil {
		return nil, err
	}

	jt.renderFinalEnvs()
	jt.removeNotFoundEnvsFromTemplate()

	if err := jt.jobModel.FromYAML([]byte(jt.template)); err != nil {
		return nil, err
	}

	return jt.jobModel.ToJobConfig(jt.finalEnvs), nil
}

func (jt *JobTemplate) removeNotFoundEnvsFromTemplate() {
	jt.template = MatchAnyEnv.ReplaceAllString(jt.template, "")
}

func (jt *JobTemplate) renderFinalEnvs() {
	for k, v := range jt.finalEnvs {
		jt.template = MatchEnv(k).ReplaceAllString(jt.template, v)
	}
}

func (jt *JobTemplate) renderTemplateData() error {
	for k, v := range jt.data {
		jt.template = MatchTemplateData(k).ReplaceAllString(jt.template, v)
	}

	return nil
}

func (jt *JobTemplate) renderDynamic() error {
	if !jt.jobModel.Env.HasVars() {
		return nil
	}

	renderers := []func() error{
		jt.getSystemEnvValues,
		jt.renderSSMEnvValues,
		jt.renderStaticEnvValues,
		jt.renderDependantEnvValues,
	}

	for _, render := range renderers {
		if err := render(); err != nil {
			return err
		}
	}

	return nil
}

func (jt *JobTemplate) isDependent(value string) bool {
	return MatchAnyEnv.MatchString(value)
}

func (jt *JobTemplate) extractReplacers(input string) []string {
	matches := MatchAnyEnv.FindAllStringSubmatch(input, -1)

	// Initialize a slice to hold the extracted variable names.
	varNames := make([]string, 0, len(matches))

	for _, match := range matches {
		// Each match is a slice where the first element is the entire match,
		// and the second element is the captured group (the variable name in this case).
		if len(match) > 1 {
			// Add the extracted variable name to the slice.
			varNames = append(varNames, match[4])
		}
	}

	return varNames
}
