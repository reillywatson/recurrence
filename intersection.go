package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Computes the set intersection of a slice of Schedules.
type Intersection []Schedule

// Implement Schedule interface.
func (self Intersection) IsOccurring(t time.Time) bool {
	for _, r := range self {
		if r.IsOccurring(t) == false {
			return false
		}
	}

	return true
}

// Implement Schedule interface.
func (self Intersection) Occurrences(t TimeRange) chan time.Time {
	return occurrencesFor(self, t)
}

// Implement Schedule interface.
func (self Intersection) NextAfter(t time.Time) (time.Time, error) {
	var zeroTime time.Time
	if len(self) == 0 {
		return zeroTime, fmt.Errorf("no more occurrences after %s", t)
	}
	current, err := self[0].NextAfter(t)
	if err != nil {
		return zeroTime, err
	}
	matchedAll := false
	for !matchedAll {
		matchedAll = true
		for _, schedule := range self {
			if !schedule.IsOccurring(current) {
				matchedAll = false
				var err error
				current, err = schedule.NextAfter(current)
				if err != nil {
					return zeroTime, fmt.Errorf("no more occurrences after %s", t)
				}
			}
		}
	}
	return current, nil
}

// Implement json.Marshaler interface.
func (self Intersection) MarshalJSON() ([]byte, error) {
	type faux Intersection
	return json.Marshal(struct {
		faux `json:"intersection"`
	}{faux: faux(self)})
}

// Implement json.Unmarshaler interface.
func (self *Intersection) UnmarshalJSON(b []byte) error {
	var mixed interface{}

	json.Unmarshal(b, &mixed)

	switch mixed.(type) {
	case []interface{}:
		for _, value := range mixed.([]interface{}) {
			bytes, _ := json.Marshal(value)
			schedule, err := ScheduleUnmarshalJSON(bytes)
			if err != nil {
				return err
			}
			*self = append(*self, schedule)
		}
	default:
		return fmt.Errorf("intersection must be a slice")
	}

	return nil
}
