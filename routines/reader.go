package routines

import (
    "log"

    "github.com/hpcloud/tail"
)

func ReadLogFile(logPath string) {
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
