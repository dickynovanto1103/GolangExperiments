package main

import (
	"errors"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestError(t *testing.T) {
	err1 := errors.New("hello")
	err2 := errors.New("hello")
	assert.Equal(t, err1, err2)
}