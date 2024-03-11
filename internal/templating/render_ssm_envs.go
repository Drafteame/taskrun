package templating

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/Drafteame/taskrun/internal/aws/ssm"
)

func (jt *JobTemplate) renderSSMEnvValues() error {
	if reflect.DeepEqual(jt.awsConfig, aws.Config{}) {
		return errMissingAWSConfig
	}

	for name, envVar := range jt.jobModel.Env.Vars {
		if envVar.Source != "ssm" {
			continue
		}

		if jt.isDependent(envVar.Key) {
			jt.dependantEnvs = append(jt.dependantEnvs, name)
			continue
		}

		if err := jt.renderSSMVar(name, envVar.Key); err != nil {
			return err
		}
	}

	return nil
}

func (jt *JobTemplate) renderSSMVar(name string, key string) error {
	if reflect.DeepEqual(jt.awsConfig, aws.Config{}) {
		return errMissingAWSConfig
	}

	val, errGet := ssm.GetParameter(key, jt.awsConfig)
	if errGet != nil {
		return errGet
	}

	jt.finalEnvs[name] = val

	return nil
}
