package templating

func (jt *JobTemplate) renderStaticEnvValues() error {
	for name, envVar := range jt.jobModel.Env.Vars {
		if envVar.Source != "" {
			continue
		}

		if jt.isDependent(envVar.Value) {
			jt.dependantEnvs = append(jt.dependantEnvs, name)
			continue
		}

		jt.finalEnvs[name] = envVar.Value
	}

	return nil
}
