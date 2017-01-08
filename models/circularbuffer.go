package models

import (
	"container/ring"
    "sync"
    "fmt"
)

// Data structure holding hits information for the last 2 minutes
type CircularBuffer struct {
    sync.Mutex
	periodHits *ring.Ring
	totalHits map[string]int
}

// Returns a new initialized circular buffer
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
    //fmt.Println("Hitting section:", h.Section)
    cb.periodHits.Value.(map[string]int)[h.Section] += 1
    cb.totalHits[h.Section] += 1
    cb.Unlock()
}

// Reset the period hits counters
func (cb *CircularBuffer) Reset() {
    cb.Lock()
    r := cb.periodHits
    for i := 0; i < r.Len(); i++ {
		r.Value = make(map[string]int)
		r = r.Next()
	}
    cb.Unlock()
}

// Display sections most hit during the last 10s and next Period
func (cb *CircularBuffer) DisplaySectionsMostHitAndNext() {
    cb.Lock()
    fmt.Println("Sections most hit during the last 10s:")
    hits := cb.periodHits.Value.(map[string]int)
    for k, v := range hits {
        fmt.Printf("\t-> %s: %v hits\n",  k, v)
    }
    fmt.Println()
    cb.periodHits.Next()
    cb.Unlock()
}
