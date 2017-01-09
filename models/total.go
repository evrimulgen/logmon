package models

// Struct holding all information since the begining of the monitoring
type Total struct {
    hits map[string]uint64
    scBytes uint64
}

// Returns a new intstantiate Total struct
func NewTotal() *Total {
    return &Total{make(map[string]uint64), 0}
}
