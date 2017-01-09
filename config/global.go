package config

import "time"

const (
	// Duration before refreshing stats
	PERIOD = 10 * time.Second
	// Number of periods taken in account in the alert logic
	NB_PERIOD = 12
)
