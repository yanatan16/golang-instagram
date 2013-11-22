package instagram

import (
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

// Sometimes location Id is a string and sometimes its an integer
type LocationId interface{}

func ParseLocationId(lid LocationId) string {
	if lid == nil {
		return ""
	}
	if slid, ok := lid.(string); ok {
		return slid
	}
	if ilid, ok := lid.(int64); ok {
		return fmt.Sprintf("%d", ilid)
	}
	return ""
}
