package workers

import (
	"context"
)

type (
	Cache interface {
		Incr(ctx context.Context, key string) error
		Get(ctx context.Context, key string) (int, bool)
	}
)
