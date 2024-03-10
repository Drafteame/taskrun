package config

type options struct {
	region    string
	profile   string
	accessKey string
	secretKey string
}

type Option func(*options)

func WithRegion(region string) Option {
	return func(o *options) {
		o.region = region
	}
}

func WithProfile(profile string) Option {
	return func(o *options) {
		o.profile = profile
	}
}

func WithAccessKey(accessKey string) Option {
	return func(o *options) {
		o.accessKey = accessKey
	}
}

func WithSecretKey(secretKey string) Option {
	return func(o *options) {
		o.secretKey = secretKey
	}
}
