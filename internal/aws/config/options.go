package config

type options struct {
	region  string
	profile string
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
