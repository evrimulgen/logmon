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
    total *Total
	alert     bool
}

// Returns a new initialized CircularBuffer
func NewCircularBuffer() *CircularBuffer {
	r := ring.New(config.NB_PERIOD)
	for i := 0; i < r.Len(); i++ {
		r.Value = NewPeriod()
		r = r.Next()
	}
	return &CircularBuffer{sync.Mutex{}, r, NewTotal(), false}
}

// Record all information of the hit into the circular buffer
func (cb *CircularBuffer) HitBy(h Hit) {
	cb.Lock()
	period := cb.periods.Value.(*Period)
	if time.Since(h.Dt) <= config.PERIOD {
		period.hits[h.Section] += 1
		period.nbHits += 1
        period.nbSCBytes += h.SCBytes
	}
	cb.total.hits[h.Section] += 1
    cb.total.scBytes += h.SCBytes
	cb.Unlock()
}

// Returns the total number of hits received by the circular buffer since the begining
func (cb *CircularBuffer) TotNbHits() uint64 {
	var totNbHits uint64
	for _, v := range cb.total.hits {
		totNbHits += v
	}
	return totNbHits
}

// Executes all monitoring tasks for a given period and launch the next one
func (cb *CircularBuffer) NextPeriod(threshold uint64) {
	cb.Lock()
	fmt.Println("##################################################################\n")
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
    fmt.Println("[INFO] Global stats:")
    fmt.Printf("\t%v hits received by the server\n", cb.TotNbHits())
    fmt.Printf("\t%v bytes sent by the server\n", cb.total.scBytes)
	period := cb.periods.Value.(*Period)
	var averageNbHits uint64
	if len(period.hits) > 0 {
		averageNbHits = period.nbHits / uint64(len(period.hits))
	}
    fmt.Println("[INFO] Period stats:")
    fmt.Printf("\t%v hits on average\n", averageNbHits)
    fmt.Printf("\t%v bytes sent by the server\n", period.nbSCBytes)
	fmt.Printf("[INFO] Sections most hit during the last %v:\n", config.PERIOD)
	for k, v := range period.hits {
		if v > averageNbHits {
			fmt.Printf("\t-> %s: %v hits (%v hits in total)\n", k, v, cb.total.hits[k])
		}
	}
	fmt.Println()
}
