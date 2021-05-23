package main

import (
	"log"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGroup(t *testing.T) {
	log.Printf("done1")
	assert.Equal(t, 1, Group())
}

func TestGroup2(t *testing.T) {
	log.Printf("done2")
	assert.Equal(t, 1, Group())
}

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		a int
		b int

		expectedSum int
	}{
		"negative": {
			a: -1,
			b: -2,

			expectedSum: -3,
		},
		"positive": {
			a: 1,
			b: 2,

			expectedSum: 3,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, test.expectedSum, Add(test.a, test.b))
		})
	}
}
