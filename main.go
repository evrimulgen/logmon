package main

import (
    "log"
    "strconv"
    "os"
    "sync"

    "github.com/gabsn/logmon/routines"
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
    go routines.ReadLogFile(logPath)
    wg.Wait()
}
