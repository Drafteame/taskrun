package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func GetParameter(param string, cfg aws.Config) (string, error) {
	client := ssm.NewFromConfig(cfg)

	out, errGet := client.GetParameter(context.Background(), &ssm.GetParameterInput{
		Name: aws.String(param),
	})

	if errGet != nil {
		return "", errGet
	}

	if out.Parameter.Value != nil {
		return *out.Parameter.Value, nil
	}

	return "", nil
}
