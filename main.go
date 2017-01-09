package main

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/gabsn/logmon/feeder"
	"github.com/gabsn/logmon/models"
	"github.com/gabsn/logmon/controller"
	"github.com/gabsn/logmon/config"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage:\n\n\tlogmon [logPath] [threshold]\n")
	}
	logPath := os.Args[1]
	_, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
    cb := models.NewCircularBuffer(config.NB_PERIOD)
	wg.Add(1)
	go feeder.ReadLogFile(logPath, cb)
    go controller.Monitor(cb)
	wg.Wait()
}
