package rl_test

import (
	"testing"
	"time"

	"github.com/mrbelka12000/rl"
)

func TestDefault(t *testing.T) {
	rl.New(
		rl.WithMethod(rl.MethodTokenBucket),
		rl.WithLimit(15),
		rl.WithTTL(2*time.Second),
	)

	var slice []int
	go func() {
		slice = append(slice, 1)
	}()
	time.Sleep(35 * time.Second)
}
