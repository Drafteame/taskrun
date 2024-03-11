package templating

import (
	"os"
)

func (jt *JobTemplate) getSystemEnvValues() error {
	for name, envVar := range jt.jobModel.Env.Vars {
		if envVar.Source != "env" {
			continue
		}

		if jt.isDependent(envVar.Key) {
			jt.dependantEnvs = append(jt.dependantEnvs, name)
			continue
		}

		jt.renderSystemEnvVar(name, envVar.Key)
	}

	return nil
}

func (jt *JobTemplate) renderSystemEnvVar(name string, key string) {
	jt.finalEnvs[name] = os.Getenv(key)
}
