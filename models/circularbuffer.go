package models

import (
	"container/ring"
    "sync"
    "fmt"

    "github.com/gabsn/logmon/config"
)

// Data structure holding hits information for the last 2 minutes
type CircularBuffer struct {
    sync.Mutex
	periodHits *ring.Ring
	totalHits map[string]uint
}

// Data structure holding all information about a given period
type Period struct {
    hits map[string]uint
    nbHits uint
}

// Returns a new intialized Period
func NewPeriod() *Period {
    return &Period{make(map[string]uint), 0}
}

// Returns a new initialized CircularBuffer
func NewCircularBuffer() *CircularBuffer {
	r := ring.New(config.NB_PERIOD)
	for i := 0; i < r.Len(); i++ {
		r.Value = NewPeriod()
		r = r.Next()
	}
	return &CircularBuffer{sync.Mutex{}, r, make(map[string]uint)}
}

// Increments the counter of hits
func (cb *CircularBuffer) HitBy(h Hit) {
    cb.Lock()
    period := cb.periodHits.Value.(*Period)
    period.hits[h.Section] += 1
    period.nbHits += 1
    cb.totalHits[h.Section] += 1
    cb.Unlock()
}

// Display sections most hit during the last 10s and next Period
func (cb *CircularBuffer) DisplayStatsAndNext() {
    cb.Lock()
    cb.displayStats()
    cb.periodHits = cb.periodHits.Next()
    cb.periodHits.Value = NewPeriod()
    cb.Unlock()
}

// Display statistics related to a given period
func (cb *CircularBuffer) displayStats() {
    fmt.Println("Sections most hit during the last 10s:")
    hits := cb.periodHits.Value.(*Period).hits
    for k, v := range hits {
        fmt.Printf("\t-> %s: %v hits\n",  k, v)
    }
    fmt.Println()
}
