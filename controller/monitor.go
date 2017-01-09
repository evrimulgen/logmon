package controller

import (
    "time"

    "github.com/gabsn/logmon/models"
    "github.com/gabsn/logmon/config"
)

func Monitor(threshold uint64, cb *models.CircularBuffer) {
    for _ = range time.Tick(config.PERIOD) {
        cb.NextPeriod(threshold)
    }
}
