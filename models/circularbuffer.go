package models

import (
	"container/ring"
	"fmt"
	"sync"
	"time"

	"github.com/gabsn/logmon/config"
)

// Data structure holding hits information for the last config.NB_PERIOD * config.PERIOD
type CircularBuffer struct {
	sync.Mutex
	periods   *ring.Ring
	totalHits map[string]uint64
	alert     bool
}

// Data structure holding all information about a given period
type Period struct {
	hits   map[string]uint64
	nbHits uint64
}

// Returns a new intialized Period
func NewPeriod() *Period {
	return &Period{make(map[string]uint64), 0}
}

// Returns a new initialized CircularBuffer
func NewCircularBuffer() *CircularBuffer {
	r := ring.New(config.NB_PERIOD)
	for i := 0; i < r.Len(); i++ {
		r.Value = NewPeriod()
		r = r.Next()
	}
	return &CircularBuffer{sync.Mutex{}, r, make(map[string]uint64), false}
}

// Increments the counter of hits
func (cb *CircularBuffer) HitBy(h Hit) {
	cb.Lock()
	period := cb.periods.Value.(*Period)
	period.hits[h.Section] += 1
	period.nbHits += 1
	cb.totalHits[h.Section] += 1
	cb.Unlock()
}

// Executes all monitoring tasks for a given period and launch the next one
func (cb *CircularBuffer) NextPeriod(threshold uint64) {
	cb.Lock()
	// Check alert with the given threshold
	cb.checkAlert(threshold)
	// Display stats related to the period
	cb.displayStats()
	// Launch the next period
	cb.periods = cb.periods.Next()
	// Initialize the next period
	cb.periods.Value = NewPeriod()
	cb.Unlock()
}

// Private method in charge of the alert logic
func (cb *CircularBuffer) checkAlert(threshold uint64) {
	var totHits uint64
	cb.periods.Do(func(x interface{}) {
		totHits += x.(*Period).nbHits
	})
	if totHits > threshold {
		cb.alert = true
		fmt.Printf("[WARN] High traffic generated an alert - hits = %v, triggered at %v\n", totHits, time.Now())
	}
	if cb.alert && totHits <= threshold {
		cb.alert = false
		fmt.Println("[WARN] Alert recovered.")
	}
}

// Display statistics related to a given period
func (cb *CircularBuffer) displayStats() {
	period := cb.periods.Value.(*Period)
	averageNbHits := period.nbHits / uint64(len(period.hits))
	fmt.Printf("[INFO] Sections most hit during the last %v (%v hits on average):\n", config.PERIOD, averageNbHits)
	for k, v := range period.hits {
		if v > averageNbHits {
			fmt.Printf("\t-> %s: %v hits\n", k, v)
		}
	}
	fmt.Println()
}
