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
		rl.WithTTL(24*time.Hour),
	)

	err := rateLimiter.Lock("suka")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("OK")
}
