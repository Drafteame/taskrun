package templating

import (
	"context"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/Drafteame/taskrun/internal/aws/secretsmanager"
	"github.com/Drafteame/taskrun/internal/models"
)

func (jt *JobTemplate) renderRemote() error {
	if jt.jobModel.Env.Remote.Type == "secretsmanager" {
		return jt.remoteSecretsManager(jt.jobModel.Env.Remote)
	}

	return nil
}

func (jt *JobTemplate) remoteSecretsManager(env models.EnvRemote) error {
	if reflect.DeepEqual(jt.awsConfig, aws.Config{}) {
		return errMissingAWSConfig
	}

	envs, err := secretsmanager.GetSecret(context.Background(), env.Key, jt.awsConfig)
	if err != nil {
		return err
	}

	for k, v := range envs {
		jt.finalEnvs[k] = v
	}

	return nil
}
