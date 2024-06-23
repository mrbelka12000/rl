package rtrl

type (
	option func(*RateLimit)
)

func WithMethod(m method) option {
	return func(rl *RateLimit) {
		rl.Method = m
	}
}
