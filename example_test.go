package rl_test

import (
	"fmt"
	"testing"

	"github.com/mrbelka12000/rl"
)

func TestDefault(t *testing.T) {
	rateLimiter := rl.New(
		rl.WithMethod(rl.MethodTokenBucket),
	)

	fmt.Printf("%+v\n", rateLimiter)
}
