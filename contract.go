package rl

type (
	locker interface {
		Lock(key string) error
	}
)
