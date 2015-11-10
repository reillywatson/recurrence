package recurrence

import (
	"encoding/json"
	"fmt"
	"time"
)

// Computes the set union of a slice of Schedules.
type Union []Schedule

// Implement Schedule interface.
func (self Union) IsOccurring(t time.Time) bool {
	for _, r := range self {
		if r.IsOccurring(t) {
			return true
		}
	}

	return false
}

// Implement Schedule interface.
func (self Union) Occurrences(t TimeRange) chan time.Time {
	return occurrencesFor(self, t)
}

// Implement Schedule interface
func (self Union) NextAfter(t time.Time) (time.Time, error) {
	var zeroTime time.Time
	if len(self) == 0 {
		return zeroTime, fmt.Errorf("no more occurrences after %s", t)
	}
	earliest := time.Time{}
	var err error
	for _, schedule := range self {
		next, schedErr := schedule.NextAfter(t)
		if err == nil && (earliest.IsZero() || next.Before(earliest)) {
			earliest = next
		}
		if earliest.IsZero() {
			err = schedErr
		} else {
			err = nil
		}
	}
	return earliest, err
}

// Implement json.Marshaler interface.
func (self Union) MarshalJSON() ([]byte, error) {
	type faux Union
	return json.Marshal(struct {
		faux `json:"union"`
	}{faux: faux(self)})
}

// Implement json.Unmarshaler interface.
func (self *Union) UnmarshalJSON(b []byte) error {
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
		return fmt.Errorf("union must be a slice")
	}

	return nil
}
