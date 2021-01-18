package uuid

import (
	"log"
	"testing"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	uuid := uuid.New()
	log.Printf("uuid: %v", uuid.String())
}
