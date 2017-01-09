package feeder

import (
	"testing"
    "reflect"

    "github.com/gabsn/logmon/models"
)

func TestParser(t *testing.T) {
	cb := models.NewCircularBuffer()
	line1 := "#Version: 1.0"
	parse(line1, cb)
	if cb.TotNbHits() != 0 || len(fields) > 0 {
        t.Error()
    }

	line2 := "#Date: 2002-05-02 17:42:15"
    parse(line2, cb)
    if cb.TotNbHits() != 0 || len(fields) > 0 {
        t.Error()
    }

	line3 := "#Fields: date time c-ip cs-username s-ip s-port cs-method cs-uri-stem sc-status cs(User-Agent)"
    fieldsExpected := []string{"date", "time", "c-ip", "cs-username", "s-ip", "s-port", "cs-method", "cs-uri-stem", "sc-status", "cs(User-Agent)"}
    parse(line3, cb)
    if !reflect.DeepEqual(fields, fieldsExpected) {
        t.Error()
    }

	line4 := "2002-05-02 17:42:15 172.22.255.255 - 172.30.255.255 80 GET /images/picture.jpg 200 Mozilla/4.0+"
    parse(line4, cb)
    if cb.TotNbHits() != 1 {
        t.Error()
    }
}
