package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/Drafteame/taskrun/internal/aws/config/credentials"
)

func Get(opts ...Option) (aws.Config, error) {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	awsOpts := make([]func(*config.LoadOptions) error, 0, 2)

	if o.region != "" {
		awsOpts = append(awsOpts, config.WithRegion(o.region))
	}

	if o.profile != "" {
		awsOpts = append(awsOpts, config.WithSharedConfigProfile(o.profile))
	}

	if o.accessKey != "" {
		provider := credentials.NewProvider(aws.Credentials{
			AccessKeyID:     o.accessKey,
			SecretAccessKey: o.secretKey,
		})

		awsOpts = append(awsOpts, config.WithCredentialsProvider(provider))
	}

	return config.LoadDefaultConfig(context.Background(), awsOpts...)
}
