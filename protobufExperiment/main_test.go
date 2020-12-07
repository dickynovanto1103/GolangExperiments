package protobufExperiment

import (
	"log"
	"testing"

	"github.com/dickynovanto1103/GolangExperiments/protobufExperiment/test"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestMarshallPairUint32_ThenUnmarshallInPairInt32(t *testing.T) {
	pairUint32 := &test.PairUint32{
		First:  proto.Uint32(1),
		Second: proto.Uint32(2),
	}

	marshalledPairUint32, err := proto.Marshal(pairUint32)
	if err != nil {
		log.Fatalf("error in proto marshall pairUint32, err: %v", err)
	}

	pairInt32 := &test.PairInt32{}
	err = proto.Unmarshal(marshalledPairUint32, pairInt32)
	if err != nil {
		log.Fatalf("error in proto unmarshall to pairInt32, err: %v", err)
	}

	expectedPairInt32 := &test.PairInt32{
		First:  proto.Int32(1),
		Second: proto.Int32(2),
	}
	assert.Equal(t, expectedPairInt32, pairInt32)
}
