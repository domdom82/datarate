package datarate

import (
	"errors"
	"strings"
	"time"
)
import "github.com/cloudfoundry/bytefmt"

type Datarate struct {
	bytes    uint64
	duration time.Duration
}

/*
Parse reads a string in the format "<number><unit>/<duration>" and produces a Datarate object.
Example: datarate.Parse("10kb/s")
*/
func Parse(s string) (*Datarate, error) {
	if strings.Count(s, "/") != 1 {
		return nil, errors.New("expected exactly one '/' in string but got: " + s)
	}
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return nil, errors.New("expected exactly one field before '/' and one after from parsing: " + s)
	}
	dataField := parts[0]
	durationField := parts[1]

	bCount, err := bytefmt.ToBytes(dataField)

	if err != nil {
		return nil, errors.New("expected data field to parse as bytes but got: " + err.Error())
	}

	//to be able to parse we inject a quantifier of 1 so "s" becomes "1s"
	dur, err := time.ParseDuration("1" + durationField)

	if err != nil {
		return nil, errors.New("expected duration field to parse as time.Duration but got: " + err.Error())
	}

	parsedRate := &Datarate{
		bytes:    bCount,
		duration: dur,
	}
	return parsedRate, nil
}
