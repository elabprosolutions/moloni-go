package models

import "time"

// Time wraps time.Time
type Time struct {
	time.Time
}

// UnmarshalJSON customizes the unmarshaling of time to handle the specific format.
func (t *Time) UnmarshalJSON(b []byte) error {
	strTime := string(b[1 : len(b)-1])

	parsedTime, err := time.Parse("2006-01-02T15:04:05-0700", strTime)
	if err != nil {
		return err
	}

	t.Time = parsedTime
	return nil
}
