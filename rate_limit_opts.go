package rl

import "time"

type (
	option func(*RateLimitBuilder)
)

// WithMethod set custom rate limit method
func WithMethod(m method) option {
	return func(rl *RateLimitBuilder) {
		rl.method = m
	}
}

// WithRedisAddr set connection to redis
func WithRedisAddr(addr string) option {
	return func(rl *RateLimitBuilder) {
		rl.redisAddr = addr
	}
}

// WithTTL set custom TTL
func WithTTL(ttl time.Duration) option {
	return func(rl *RateLimitBuilder) {
		rl.ttl = ttl
	}
}

// WithLimit set custom limit
func WithLimit(limit uint) option {
	return func(rl *RateLimitBuilder) {
		rl.limit = limit
	}
}
