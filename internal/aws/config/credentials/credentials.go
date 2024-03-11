package credentials

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
)

// Credentials is a compound structure that holds AWS access credentials.
type Credentials struct {
	// AWS configured region
	Region string

	// AWS Access key ID
	AccessKeyID string

	// AWS Secret Access Key
	SecretAccessKey string

	// AWS Session Token
	SessionToken string

	// Source of the credentials
	Source string

	// States if the credentials can expire or not.
	CanExpire bool

	// The time the credentials will expire at. Should be ignored if CanExpire
	// is false.
	Expires time.Time
}

// Retrieve obtain the current configuration of AWS credentials.
func Retrieve(ctx context.Context) (*Credentials, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	credentials, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, err
	}

	return &Credentials{
		Region:          cfg.Region,
		AccessKeyID:     credentials.AccessKeyID,
		SecretAccessKey: credentials.SecretAccessKey,
		SessionToken:    credentials.SessionToken,
		Source:          credentials.Source,
		CanExpire:       credentials.CanExpire,
		Expires:         credentials.Expires,
	}, nil
}
