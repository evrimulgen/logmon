package controller

import (
	"time"

	"github.com/gabsn/logmon/config"
	"github.com/gabsn/logmon/models"
)

func Monitor(threshold uint64, cb *models.CircularBuffer) {
	for _ = range time.Tick(config.PERIOD) {
		cb.NextPeriod(threshold)
	}
}
