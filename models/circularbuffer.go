package models

import (
	"container/ring"
    "sync"
)

type CircularBuffer struct {
    sync.Mutex
	PeriodHits *ring.Ring
	TotalHits map[string]int
}

func NewCircularBuffer(nbPeriod int) *CircularBuffer {
	r := ring.New(nbPeriod)
	for i := 0; i < r.Len(); i++ {
		r.Value = make(map[string]int)
		r = r.Next()
	}
	return &CircularBuffer{sync.Mutex{}, r, make(map[string]int)}
}

// Increments the counter of hits
func (cb *CircularBuffer) HitBy(h Hit) {
    cb.Lock()
    cb.PeriodHits.Value.(map[string]int)[h.Section] += 1
    cb.TotalHits[h.Section] += 1
    cb.Unlock()
}

// Reset the period hits counters
func (cb *CircularBuffer) Reset() {
    cb.Lock()
    r := cb.PeriodHits
    for i := 0; i < r.Len(); i++ {
		r.Value = make(map[string]int)
		r = r.Next()
	}
    cb.Unlock()
}
