package rtrl

type RateLimit struct {
	Method method
}

func New(opts ...option) *RateLimit {
	rl := &RateLimit{
		Method: MethodTokenBucket,
	}

	for _, opt := range opts {
		opt(rl)
	}

	return rl
}
