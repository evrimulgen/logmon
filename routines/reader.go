package routines

import (
    "log"
	"regexp"
    "strings"
	"errors"
    "fmt"
    "time"

    "github.com/hpcloud/tail"
)

var (
    fields []string
    fieldsRE = regexp.MustCompile(`#Fields: (\S+\s?)+`)
)

type Section struct {
    time time.Time
    section string
}

func NewSection(l map[string]string) Section {
	return Section{
		getTime(l["time"]),
		getSection(l["cs-uri-stem"]),
	}
}

func getTime(t string) time.Time {
	new_time, err := time.Parse("15:04:05", t)
	if err != nil {
		panic(err)
	}
	hour, minute, second := new_time.Clock()
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, hour, minute, second, 0, time.UTC)
}

func getSection(uri string) string {
	sectionSplit := strings.Split(uri, "/")
	if len(sectionSplit) < 2 {
		panic(errors.New("Invalid resource format"))
	}
	return strings.Join(sectionSplit[:2], "/")
}

func ReadLogFile(logPath string) {
    t, err := tail.TailFile(logPath, tail.Config{
		Follow:   true,
		ReOpen:   true,
	})
	if err != nil {
        log.Fatalln(err)
	}
	for line := range t.Lines {
        parseLine(line.Text)
	}
}

func parseLine(line string) {
    if fields == nil {
        parseHeader(line)
        return
    } else {
		section := NewSection(parseToMap(line))
        fmt.Println(section.time.String(), section.section)
    }
}

func parseHeader(line string) {
	if fieldsRE.MatchString(line) {
		fields = strings.Split(line, " ")
        fields = fields[1:]
	}
}

func parseToMap(line string) map[string]string {
    l := make(map[string]string)
    data := strings.Split(line, " ")
    for k, v := range(fields) {
        if v != "" {
            l[v] = data[k]
        }
    }
	return l
}
