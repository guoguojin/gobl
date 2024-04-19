package cache

import "time"

type options struct {
	// expiry is the default expiry for items in the cache
	expiry time.Duration
}

type Option func(*options)

func WithExpiry(expiry time.Duration) Option {
	return func(o *options) {
		o.expiry = expiry
	}
}

func defaultOptions() options {
	return options{
		expiry: NoExpiry,
	}
}
