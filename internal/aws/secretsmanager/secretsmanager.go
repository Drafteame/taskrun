package secretsmanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// MustGetSecret retrieves the secret value for a secret name from the secrets manager.
func MustGetSecret(secretName string, cfg aws.Config) map[string]string {
	client := secretsmanager.NewFromConfig(cfg)

	result := make(map[string]string)

	getSecretValueInput := secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	getSecretValueOutput, err := client.GetSecretValue(context.Background(), &getSecretValueInput)
	if err != nil {
		panic(err)
	}

	if getSecretValueOutput.SecretString == nil {
		panic(err)
	}

	if err = json.Unmarshal([]byte(*getSecretValueOutput.SecretString), &result); err != nil {
		panic(err)
	}

	return result
}
