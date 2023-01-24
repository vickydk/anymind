package utils

import (
	"math/rand"
	"time"
)

// Policy implements back off policy, random delay in millis
type Policy struct {
	Millis []int
}

// Default is back off policy up to 5 seconds
var Default = Policy{
	[]int{0, 1000, 1000, 3000, 3000, 5000, 5000, 7000, 9000, 10000},
}

// Duration return time duration
func (b Policy) Duration(n int) time.Duration {
	if n >= len(b.Millis) {
		n = len(b.Millis) - 1
	}

	return time.Duration(jitter(b.Millis[n])) * time.Millisecond
}

// [0.5 * millis .. 1.5 * millis]
func jitter(millis int) int {
	if millis == 0 {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return millis/2 + rand.Intn(millis)
}
