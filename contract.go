package rl

import (
	"context"
	"time"
)

type (
	locker interface {
		Incr(ctx context.Context, key string) error
		Expire(ctx context.Context, key string, ttl time.Duration) error
		Get(ctx context.Context, key string) (string, error)
	}
)
