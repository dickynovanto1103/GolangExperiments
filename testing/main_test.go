package main

import (
	"log"
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMain(m *testing.M) {
	log.Printf("before test main")
	os.Exit(m.Run())
	log.Printf("done test main")
}

func TestGroup(t *testing.T) {
	log.Printf("done1")
	assert.Equal(t, 1, Group())
}

func TestGroup2(t *testing.T) {
	log.Printf("done2")
	assert.Equal(t, 1, Group())
}
