package pkg

import (
	"encoding/json"
	"time"
)

type TimeZone string

const DateTimeFormatStandard = time.RFC3339

const (
	TimeZoneLondon TimeZone = "Europe/London"
)

type DateTime struct {
	time.Time
}

func NewDateTime() DateTime {
	// Default to London's time (centre of my universe)
	timezoneLocation, _ := time.LoadLocation(string(TimeZoneLondon))

	return DateTime{time.Now().In(timezoneLocation)}
}

func (receiver *DateTime) UnmarshalJSON(b []byte) error {
	var parsedTime time.Time
	err := json.Unmarshal(b, &parsedTime)
	if err != nil {
		return err
	}

	receiver.Time = parsedTime

	return nil
}

func (receiver *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(receiver.Format(DateTimeFormatStandard)), nil
}

func (receiver *DateTime) String() string {
	return receiver.Format(DateTimeFormatStandard)
}
