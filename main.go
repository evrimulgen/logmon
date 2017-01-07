package main

import (
    "log"
    "strconv"
    "os"

    "github.com/hpcloud/tail"
)

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

    t, err := tail.TailFile(logPath, tail.Config{
		Follow:   true,
		ReOpen:   true,
	})
	if err != nil {
        log.Fatalln(err)
	}
	for line := range t.Lines {
        println(line.Text)
	}
}
