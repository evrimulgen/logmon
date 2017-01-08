package monitor

import (
    "time"

    "github.com/gabsn/logmon/models"
)

func Supervise(period time.Duration, cb *models.CircularBuffer) {
    for _ = range time.Tick(period) {
        cb.DisplaySectionsMostHitAndNext()
    }
}
