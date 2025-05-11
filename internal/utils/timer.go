package utils

import "time"

type Timer struct {
	start    time.Time
	duration time.Duration
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{start: time.Now(), duration: duration}
}

func (t *Timer) Done() bool {
	return time.Since(t.start) >= t.duration
}

func (t *Timer) Elapsed() time.Duration {
	return time.Since(t.start)
}

func (t *Timer) Reset() {
	t.start = time.Now()
}

func (t *Timer) Finish() {
	t.start = time.Now().Add(-t.duration)
}
