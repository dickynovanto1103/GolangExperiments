package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func generateError() error {
	return errors.Wrap(errors.New("new error"), "got error")
}

func main() {
	err := errors.Wrap(generateError(), "error in main")
	err1 := fmt.Errorf("error new: %w", err)
	//in order for unwrap(err1) returning error wrapped in err1, previously when creating err1, it must use fmt.Errorf and then use verb %w to identify the error wrapped
	// if we use other verbs, like %v, the errors.Unwrap(err1) will return nil
	err2 := errors.Unwrap(err1)

	log.Println("err:", err2)
}
