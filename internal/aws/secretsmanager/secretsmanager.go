package secretsmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// GetSecret retrieves the secret value for a secret name from the secrets manager.
func GetSecret(ctx context.Context, secretName string, cfg aws.Config) (map[string]string, error) {
	client := secretsmanager.NewFromConfig(cfg)

	result := make(map[string]any)

	getSecretValueInput := secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	getSecretValueOutput, err := client.GetSecretValue(ctx, &getSecretValueInput)
	if err != nil {
		return nil, err
	}

	if getSecretValueOutput.SecretString == nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(*getSecretValueOutput.SecretString), &result); err != nil {
		return nil, err
	}

	// cast the result to string
	resString := make(map[string]string)

	for k, v := range result {
		resString[k] = fmt.Sprintf("%v", v)
	}

	return resString, nil
}
