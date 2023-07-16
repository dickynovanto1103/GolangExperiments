package go_test_caching_experiment

import (
	"fmt"
	"os"
)

func Hello(str string) string {
	debug := os.Getenv("DEBUG_HELLO")
	if len(debug) > 0 {
		return fmt.Sprintf("debug: %v", debug)
	}
	if len(str) > 0 {
		return fmt.Sprintf("Hello %v", str)
	}

	return "Hello"
}
