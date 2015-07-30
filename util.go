package protorpc

import (
	"math/rand"
	"time"
)

const (
	// how long to wait after the first failure before retrying
	baseDelay = 1.0 * time.Second
	// upper bound of backoff delay
	maxDelay = 120 * time.Second
	// backoff increases by this factor on each retry
	backoffFactor = 1.6
	// backoff is randomized downwards by this factor
	backoffJitter = 0.2
)

func backoff(retries int) (t time.Duration) {
	if retries == 0 {
		return baseDelay
	}
	backoff, max := float64(baseDelay), float64(maxDelay)
	for backoff < max && retries > 0 {
		backoff *= backoffFactor
		retries--
	}
	if backoff > max {
		backoff = max
	}
	// Randomize backoff delays so that if a cluster of requests start at
	// the same time, they won't operate in lockstep.
	backoff *= 1 + backoffJitter*(rand.Float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return time.Duration(backoff)
}
