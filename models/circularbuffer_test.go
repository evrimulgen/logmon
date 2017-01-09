package models

import (
	"testing"
)

func TestAlertLogic(t *testing.T) {
	cb := NewCircularBuffer()
	cb.checkAlert(20)
	// No hit yet, so no alert should be generated
	if cb.alert == true {
		t.Error()
	}

	cb.periods.Value = &Period{make(map[string]uint64), 30}
	cb.checkAlert(20)
	// An alert should be generated here because a period registered
	// more hits than the threshold
	if cb.alert == false {
		t.Error()
	}

	cb.periods = cb.periods.Next()
	cb.periods.Value = &Period{make(map[string]uint64), 10}
	cb.checkAlert(20)
	// The alert should still be there because the total number of hits
	// is now 40, which is above the threshold
	if cb.alert == false {
		t.Error()
	}

	cb.periods = cb.periods.Prev()
	cb.periods.Value = &Period{make(map[string]uint64), 5}
	cb.checkAlert(20)
	// The alert should be stopped because the total number of hits
	// is now of 15, which is below the threshold
	if cb.alert == true {
		t.Error()
	}
}
