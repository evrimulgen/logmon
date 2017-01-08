package feeder

import (
	"log"

	"github.com/hpcloud/tail"
)

// Goroutine that consumes log.txt lines and send them to the parser
func ReadLogFile(logPath string) {
	t, err := tail.TailFile(logPath, tail.Config{
		Follow: true,
		ReOpen: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	for line := range t.Lines {
		parse(line.Text)
	}
}
