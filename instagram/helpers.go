package instagram

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type StringUnixTime string

func (s StringUnixTime) Time() (t time.Time, err error) {
	unix, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return
	}

	t = time.Unix(unix, 0)
	return
}

// UnmarshalJSON is implemented on Location because the ID can be
// returned as a string or an int, and we would like to always have
// it as a string on the Location struct.
func (l *Location) UnmarshalJSON(b []byte) error {
	temp := &struct {
		ID        interface{}
		Name      string
		Latitude  float64
		Longitude float64
	}{}
	if err := json.Unmarshal(b, temp); err != nil {
		return err
	}
	switch id := temp.ID.(type) {
	case int:
		l.Id = strconv.Itoa(id)
	case string:
		l.Id = id
	default:
		return fmt.Errorf("unknown type received for location id: %T", id)
	}
	l.Latitude = temp.Latitude
	l.Longitude = temp.Longitude
	l.Name = temp.Name
	return nil
}
