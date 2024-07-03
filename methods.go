package rl

type (
	method string
)

const (
	MethodTokenBucket          method = "token_bucket"
	MethodLeakyBucket          method = "leaky_bucket"
	MethodFixedWindow          method = "fixed_window"
	MethodSlidingWindow        method = "sliding_window"
	MethodSlidingWindowCounter method = "sliding_window_counter"
)
