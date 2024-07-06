package rl

import (
	"time"

	"github.com/mrbelka12000/rl/cache/inmem"
	"github.com/mrbelka12000/rl/cache/redis"
	"github.com/mrbelka12000/rl/errs"
	rpkg "github.com/mrbelka12000/rl/pkg/redis"
	"github.com/mrbelka12000/rl/workers"
	"github.com/mrbelka12000/rl/workers/fixed_window"
	"github.com/mrbelka12000/rl/workers/leacky_bucket"
	"github.com/mrbelka12000/rl/workers/sliding_window"
	"github.com/mrbelka12000/rl/workers/sliding_window_counter"
	"github.com/mrbelka12000/rl/workers/token_bucket"
)

const (
	defaultTTL   = 5 * time.Second
	defaultLimit = 10
)

type (
	// RateLimitBuilder for build all connections
	RateLimitBuilder struct {
		method    Method
		redisAddr string
		limit     uint
		ttl       time.Duration
	}

	// RateLimiter ...
	RateLimiter struct {
		locker
	}
)

// New entrypoint that creates instance for rate limiter
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

	var w workers.Cache
	if rlb.redisAddr == "" {
		w = inmem.New(rlb.ttl)
	} else {
		client, err := rpkg.GetConnection(rlb.redisAddr)
		if err != nil {
			panic(err)
		}
		w = redis.New(client, rlb.ttl)
	}

	rl := &RateLimiter{
		locker: chooseWorker(w, rlb.method, rlb.limit),
	}

	return rl
}

func chooseWorker(w workers.Cache, m Method, limit uint) locker {
	switch m {
	case MethodFixedWindow:
		return fixed_window.New(w, limit)

	case MethodTokenBucket:
		return token_bucket.New(w, limit)

	case MethodSlidingWindowCounter:
		return sliding_window_counter.New(w, limit)

	case MethodSlidingWindow:
		return sliding_window.New(w, limit)

	case MethodLeakyBucket:
		return leacky_bucket.New(w, limit)

	default:
		panic(errs.ErrMethodUndefined)
	}
}
