package templating

import "github.com/Drafteame/taskrun/internal/models"

func (jt *JobTemplate) renderDependantEnvValues() error {
	envVars := make(map[string]models.EnvVar)

	for _, name := range jt.dependantEnvs {
		envVars[name] = jt.jobModel.Env.Vars[name]
	}

	for name, envVar := range envVars {
		if err := jt.renderDependantEnv(name, envVar); err != nil {
			return err
		}
	}

	return nil
}

func (jt *JobTemplate) renderDependantEnv(name string, env models.EnvVar) error {
	switch env.Source {
	case "ssm":
		return jt.renderDependantSSMVar(name, env)
	case "env":
		return jt.renderDependantEnvVar(name, env)
	default:
		return jt.renderDependantStaticVar(name, env)
	}
}

func (jt *JobTemplate) renderDependantSSMVar(name string, env models.EnvVar) error {
	return jt.renderDependant(name, env.Key, func(n, v string) error {
		return jt.renderSSMVar(n, v)
	})
}

func (jt *JobTemplate) renderDependantEnvVar(name string, env models.EnvVar) error {
	return jt.renderDependant(name, env.Key, func(n, v string) error {
		jt.renderSystemEnvVar(n, v)
		return nil
	})
}

func (jt *JobTemplate) renderDependantStaticVar(name string, env models.EnvVar) error {
	return jt.renderDependant(name, env.Value, func(n, v string) error {
		jt.finalEnvs[n] = v
		return nil
	})
}

func (jt *JobTemplate) renderDependant(name, value string, callback func(string, string) error) error {
	dependencies := jt.extractReplacers(value)

	for _, dep := range dependencies {
		final, ok := jt.finalEnvs[dep]
		if ok {
			value = MatchEnv(dep).ReplaceAllString(value, final)
		}

		envVar, exists := jt.jobModel.Env.Vars[dep]
		if !exists {
			value = MatchEnv(dep).ReplaceAllString(value, "")
		}

		if err := jt.renderDependantEnv(dep, envVar); err != nil {
			return err
		}
	}

	return callback(name, value)
}
