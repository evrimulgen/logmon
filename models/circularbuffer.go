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
