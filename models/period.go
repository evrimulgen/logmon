package models

// Data structure holding all information about a given period
type Period struct {
	hits      map[string]uint64
	nbHits    uint64
	nbSCBytes uint64
}

// Returns a new intialized Period
func NewPeriod() *Period {
	return &Period{make(map[string]uint64), 0, 0}
}
