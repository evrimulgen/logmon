package controller

import (
    "time"

    "github.com/gabsn/logmon/models"
    "github.com/gabsn/logmon/config"
)

func Monitor(cb *models.CircularBuffer) {
    var period_counter uint
    for _ = range time.Tick(config.PERIOD) {
        period_counter += 1
        cb.DisplaySectionsMostHitAndNext(&period_counter)
    }
}
