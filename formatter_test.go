package ylog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLineFormatter(t *testing.T) {
	record := Record{
		Time:  time.Now(),
		Msg:   "test info",
		Level: InfoLevel,
		Fields: Fields{
			"testkey": "testvalue",
		},
	}

	line := &LineFormatter{}
	str, err := line.Format(record)
	assert.Nil(t, err)
	assert.Contains(t, str, "test info")
	assert.Contains(t, str, "testvalue")
}

func TestJsonFormat(t *testing.T) {
	record := Record{
		Time:  time.Now(),
		Msg:   "test info",
		Level: InfoLevel,
		Fields: Fields{
			"testkey": "testvalue",
		},
	}

	json := &JsonFormatter{}
	str, err := json.Format(record)
	assert.Nil(t, err)
	assert.Contains(t, str, "\"msg\":\"test info\"")
	assert.Contains(t, str, "\"testkey\":\"testvalue\"")
}
