package credentials

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Provider struct {
	credentials aws.Credentials
}

func NewProvider(c aws.Credentials) Provider {
	return Provider{
		credentials: c,
	}
}

func (cp Provider) Retrieve(_ context.Context) (aws.Credentials, error) {
	return cp.credentials, nil
}
