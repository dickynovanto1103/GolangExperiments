package go_test_caching_experiment

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func genNumChars(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += "a"
	}
	return res
}

func TestHello(t *testing.T) {
	debug := os.Getenv("DEBUG_HELLO")
	if len(debug) > 0 {
		assert.Equal(t, fmt.Sprintf("debug: %v", debug), Hello("adsff"))
		return
	}

	mapp := make(map[string]struct{})
	mapp["dicky"] = struct{}{}
	msg := Hello(genNumChars(100000))
	assert.Equal(t, fmt.Sprintf("Hello %v", genNumChars(100000)), msg)
}
