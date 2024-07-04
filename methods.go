package rl

type (
	Method string
)

const (
	MethodTokenBucket          Method = "token_bucket"
	MethodLeakyBucket          Method = "leaky_bucket"
	MethodFixedWindow          Method = "fixed_window"
	MethodSlidingWindow        Method = "sliding_window"
	MethodSlidingWindowCounter Method = "sliding_window_counter"
)
