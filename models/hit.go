package models

import (
    "time"
    "fmt"
)

type Hit struct {
    Dt time.Time
    Section string
}

func (h *Hit) Hits(cb *CircularBuffer) {
    fmt.Printf("Section %s hit at %v\n", h.Section, h.Dt)
}
