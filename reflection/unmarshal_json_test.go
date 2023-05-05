package main

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		objInput    Person
		objExpected Person
	}{
		"person": {
			objInput: Person{
				Name: "dicky",
				Age:  22,
			},
			objExpected: Person{
				Name: "dicky",
				Age:  22,
			},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			bytes, err := json.Marshal(test.objInput)
			if err != nil {
				log.Panicf("error in json marshal, err: %v", err)
			}

			log.Printf("bytes: %s", bytes)

			objRecv := &Person{}
			UnmarshalJSON(bytes, objRecv)
			assert.Equal(t, test.objExpected, *objRecv)
		})
	}
}
