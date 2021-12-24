package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	num := -1
	assert.Equal(t, Negative, DetermineType(num))
}
