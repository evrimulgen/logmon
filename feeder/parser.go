package feeder

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gabsn/logmon/models"
)

type Hit models.Hit

var fields []string

var (
	fieldsRE = regexp.MustCompile(`#Fields: (\S+\s?)+`)
	dateRE   = regexp.MustCompile(`date`)
	timeRE   = regexp.MustCompile(`time`)
	uriRE    = regexp.MustCompile(`uri`)
)

// Parse a line into a Hit struct and Hits CircularBuffer
func parse(line string) {
	if fields == nil {
		parseHeader(line)
	} else {
		hit, err := parseToHit(line)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(hit.Dt.String(), hit.Section)
	}
}

// Parse log header to know what fields to take in account
func parseHeader(line string) {
	if fieldsRE.MatchString(line) {
		fields = strings.Split(line, " ")
		fields = fields[1:]
	}
}

// Parse a log line into the corresponding Hit struct
func parseToHit(line string) (Hit, error) {
	var date, time, uri string
	hitFields := strings.Split(line, " ")
	for k, v := range fields {
		switch {
		case dateRE.MatchString(v):
			date = hitFields[k]
		case timeRE.MatchString(v):
			time = hitFields[k]
		case uriRE.MatchString(v):
			uri = hitFields[k]
		}
	}
	dt, err := getDateTime(date, time)
	if err != nil {
		return Hit{}, err
	}
	section, err := getSection(uri)
	if err != nil {
		return Hit{}, err
	}
	return Hit{dt, section}, nil
}

// Build a time.Time object from date and time fields extracted
func getDateTime(d, t string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		date = time.Now()
	}
	timee, er := time.Parse("15:04:05", t)
	if er != nil {
		return date, er
	}
	hour, minute, second := timee.Clock()
	year, month, day := date.Date()
	return time.Date(year, month, day, hour, minute, second, 0, time.UTC), nil
}

// Check if the resource is valid and return the section part
func getSection(uri string) (string, error) {
	sectionSplit := strings.Split(uri, "/")
	if len(sectionSplit) < 2 {
		return "", errors.New("Invalid resource format")
	}
	return strings.Join(sectionSplit[:2], "/"), nil
}
