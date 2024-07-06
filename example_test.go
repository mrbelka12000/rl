package rl_test

import (
	"testing"
	"time"

	"github.com/mrbelka12000/rl"
)

func TestDefault(t *testing.T) {
	rateLimiter := rl.New(
		rl.WithMethod(rl.MethodTokenBucket),
		rl.WithLimit(15),
		rl.WithTTL(2*time.Second),
	)

	rateLimiter.Lock("test")
}
