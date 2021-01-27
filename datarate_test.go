package datarate

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestParse(t *testing.T) {

	table := map[string]bool{
		"10KB/s":    true,
		"10kb/s":    false,
		"10kib/s":   false,
		"3MB/s":     true,
		"kb/s":      false,
		"bad":       false,
		"123 mb//s": false,
		"1.5MB/h":   true,
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

func TestValues(t *testing.T) {

	oneMegabytePerSecond, err := Parse("1MB/s")
	assert.NoError(t, err)
	assert.Equal(t, 1e6, oneMegabytePerSecond.BytesPerSecond())

	oneGigabytePerHour, err := Parse("1GB/h")
	assert.NoError(t, err)
	assert.LessOrEqual(t, 277777.7, oneGigabytePerHour.BytesPerSecond())

	tenBytesPerSecond, err := Parse("10B/s")
	assert.NoError(t, err)
	assert.Equal(t, 10.0, tenBytesPerSecond.BytesPerSecond())

}

func TestYamlMarshal(t *testing.T) {

	type YamlDocument struct {
		Rate *Datarate `yaml:"rate,omitempty"`
	}

	input := []byte("rate: '1GB/s'")
	var output YamlDocument

	err := yaml.Unmarshal(input, &output)
	assert.NoError(t, err)
	assert.NotNil(t, output.Rate)
	assert.Equal(t, 1.0, output.Rate.GigabytesPerSecond())
}
