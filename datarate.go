package datarate

import (
	"errors"
	"github.com/alecthomas/units"
	"github.com/martinlindhe/unit"
	"strings"
	"time"
)

/*
Parse reads a string in the format "<number><unit>/<duration>" and produces a github.com/martinlindhe/unit.Datarate object.
Since data rates are usually metric (base 10), we only support those. eg. megabit/s, not mebibit/s
Example: datarate.Parse("10KB/s")  // 10 kilobyte per second
Bad Example: data.Parse("10KiB/s") // error: kibibytes not supported
*/
func Parse(s string) (unit.Datarate, error) {
	if strings.Count(s, "/") != 1 {
		return 0, errors.New("expected exactly one '/' in string but got: " + s)
	}
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return 0, errors.New("expected exactly one field before '/' and one after from parsing: " + s)
	}
	dataField := parts[0]
	durationField := parts[1]

	byteCount, err := units.ParseMetricBytes(dataField)

	if err != nil {
		return 0, errors.New("expected data field to parse as bytes but got: " + err.Error())
	}

	//we need bits per second
	bitCount := byteCount * 8

	//to be able to parse duration we inject a quantifier of 1 so "s" becomes "1s"
	dur, err := time.ParseDuration("1" + durationField)

	if err != nil {
		return 0, errors.New("expected duration field to parse as time.Duration but got: " + err.Error())
	}

	//duration is in nanoseconds so we need to convert to seconds
	seconds := dur / time.Second

	parsedRate := float64(bitCount) / float64(seconds)

	return unit.Datarate(parsedRate), nil
}
