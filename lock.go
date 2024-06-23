package rtrl

func (rl *RateLimit) Lock(key string) error {
	switch rl.Method {
	case MethodTokenBucket:
		return rl.tbLock(key)
	case MethodFixedWindow:
		return rl.fwLock(key)
	case MethodLeakyBucket:
		return rl.lbLock(key)
	case MethodSlidingWindow:
		return rl.swLock(key)
	case MethodSlidingWindowCounter:
		return rl.swcLock(key)
	default:
		return ErrMethodUndefined
	}
}
