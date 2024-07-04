package rl

import (
	"time"

	rpkg "github.com/mrbelka12000/rl/pkg/redis"
	"github.com/mrbelka12000/rl/store/inmem"
	"github.com/mrbelka12000/rl/store/redis"
)

const (
	defaultTTL   = 5 * time.Second
	defaultLimit = 10
)

type (
	// RateLimitBuilder for build all connections
	RateLimitBuilder struct {
		method    method
		redisAddr string
		limit     uint
		ttl       time.Duration
	}

	// RateLimiter implement all methods
	RateLimiter struct {
		method method
		limit  uint
		ttl    time.Duration
		Locker
	}
)

func New(opts ...option) *RateLimiter {
	// default settings for rate limiter
	rlb := &RateLimitBuilder{
		method: MethodFixedWindow,
		limit:  defaultLimit,
		ttl:    defaultTTL,
	}

	for _, opt := range opts {
		opt(rlb)
	}

	rl := &RateLimiter{
		method: rlb.method,
		limit:  rlb.limit,
		ttl:    rlb.ttl,
	}

	if rlb.redisAddr == "" {
		rl.Locker = inmem.New()
	} else {
		redisCLI, err := rpkg.GetConnection(rlb.redisAddr)
		if err != nil {
			panic(err)
		}
		rl.Locker = redis.New(redisCLI)
	}

	return rl
}
