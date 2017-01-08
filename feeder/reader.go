package feeder

import (
	"log"

	"github.com/hpcloud/tail"
    "github.com/gabsn/logmon/models"
)

// Goroutine that consumes log.txt lines and send them to the parser
func ReadLogFile(logPath string, cb *models.CircularBuffer) {
	t, err := tail.TailFile(logPath, tail.Config{
		Follow: true,
		ReOpen: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	for line := range t.Lines {
		parse(line.Text, cb)
	}
}
