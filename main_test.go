package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysPasses(t *testing.T) {
	assert.True(t, true, "True is always true")
}
