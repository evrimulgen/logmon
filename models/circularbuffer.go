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
	totalHits map[string]uint
}

// Returns a new initialized circular buffer
func NewCircularBuffer(nbPeriod int) *CircularBuffer {
	r := ring.New(nbPeriod)
	for i := 0; i < r.Len(); i++ {
		r.Value = make(map[string]uint)
		r = r.Next()
	}
	return &CircularBuffer{sync.Mutex{}, r, make(map[string]uint)}
}

// Increments the counter of hits
func (cb *CircularBuffer) HitBy(h Hit) {
    cb.Lock()
    //fmt.Println("Hitting section:", h.Section)
    cb.periodHits.Value.(map[string]uint)[h.Section] += 1
    cb.totalHits[h.Section] += 1
    cb.Unlock()
}

// Display sections most hit during the last 10s and next Period
func (cb *CircularBuffer) DisplayStatsAndNext() {
    cb.Lock()
    cb.displayStats()
    cb.periodHits = cb.periodHits.Next()
    cb.periodHits.Value = make(map[string]uint)
    cb.Unlock()
}

func (cb *CircularBuffer) displayStats() {
    fmt.Println("Sections most hit during the last 10s:")
    hits := cb.periodHits.Value.(map[string]uint)
    for k, v := range hits {
        fmt.Printf("\t-> %s: %v hits\n",  k, v)
    }
    fmt.Println()
}

// Reset the period hits counters (no mutexes because it's a private method)
func (cb *CircularBuffer) reset() {
    r := cb.periodHits
    for i := 0; i < r.Len(); i++ {
		r.Value = make(map[string]uint)
		r = r.Next()
	}
}
