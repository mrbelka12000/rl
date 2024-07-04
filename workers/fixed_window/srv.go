package fixed_window

import (
	"time"

	"github.com/mrbelka12000/rl/workers"
)

type FixedWindow struct {
	worker workers.Cache
	limit  int
	ttl    time.Duration
}

func New(w workers.Cache, limit uint, ttl time.Duration) *FixedWindow {
	return &FixedWindow{}
}

func (f FixedWindow) Lock(key string) error {
	return nil
}
