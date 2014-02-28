package main

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {

	assert.Equal(t, int64(5000), getPort())

	os.Setenv("PORT", "80")
	assert.Equal(t, int64(80), getPort())
}
