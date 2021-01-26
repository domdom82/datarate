package datarate

import (
	"testing"
)

func TestParse(t *testing.T) {

	table := map[string]bool{
		"10kb/s":    true,
		"3mb/s":     true,
		"kb/s":      false,
		"bad":       false,
		"123 mb//s": false,
		"1.5M/h":    true,
	}

	for k, v := range table {
		rate, err := Parse(k)
		if err != nil && v == true {
			t.Errorf("expected %s to parse but got: %v", k, err)
		}
		if err == nil && v == false {
			t.Errorf("expected %s to throw error but got: %v", k, rate)
		}
	}

}
