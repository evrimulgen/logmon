package main

import (
	"log"
	"os"
	"strconv"
	"sync"
    "time"

	"github.com/gabsn/logmon/feeder"
	"github.com/gabsn/logmon/models"
	"github.com/gabsn/logmon/monitor"
)

var wg sync.WaitGroup

const (
    period = 10 * time.Second
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage:\n\n\tlogmon [logPath] [threshold]\n")
	}
	logPath := os.Args[1]
	_, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
    cb := models.NewCircularBuffer(12)
	wg.Add(1)
	go feeder.ReadLogFile(logPath, cb)
    go monitor.Supervise(period, cb)
	wg.Wait()
}
