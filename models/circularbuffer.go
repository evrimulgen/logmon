package models

import (
	"container/ring"
    "sync"
)

type CircularBuffer struct {
    sync.Mutex
	TotalHits map[string]int
	PeriodHits *ring.Ring
}

func NewCircularBuffer(nbPeriod int) *CircularBuffer {
	r := ring.New(nbPeriod)
	for i := 0; i < r.Len(); i++ {
		r.Value = make(map[string]int)
		r = r.Next()
	}
	return &CircularBuffer{sync.Mutex{}, make(map[string]int), r}
}

// Increments the counter of hits
func (cb *CircularBuffer) HitBy(h Hit) {
    cb.Lock()
    PerdiodHits[h.Section] += 1
    TotalHits[h.Section] += 1
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
