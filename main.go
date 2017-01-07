package main

import (
    "log"
    "strconv"
    "os"
)

func usageAndQuit() {
   log.Fatalln("Usage:\n\n\tlogmon [logPath] [threshold]\n")
}

func main() {
    if len(os.Args) != 3 {
        usageAndQuit()
    }
    logPath := os.Args[1]
    threshold, err := strconv.Atoi(os.Args[2])
    if err != nil {
        log.Fatalln(err)
    }
}
