package main

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/gabsn/logmon/feeder"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage:\n\n\tlogmon [logPath] [threshold]\n")
	}
	logPath := os.Args[1]
	threshold, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	println(threshold)

	wg.Add(1)
	go feeder.ReadLogFile(logPath)
	wg.Wait()
}
