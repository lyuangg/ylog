package ylog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	levels := []string{
		"debug",
		"info",
		"warning",
		"error",
		"unknow",
		"unknow",
	}
	for k, v := range levels {
		level := Level(k)
		assert.Equal(t, v, level.String(), "level name err")
	}
}

func TestParseLevel(t *testing.T) {
	levels := []string{
		"debug",
		"info",
		"warning",
		"error",
	}
	for _, v := range levels {
		_, err := ParseLevel(v)
		assert.Nil(t, err)
	}

	errlevel, err := ParseLevel("unknow")
	assert.NotNil(t, err)
	assert.Equal(t, errlevel, ErrorLevel)
}
