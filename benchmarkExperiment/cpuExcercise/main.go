package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

func readByte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Panicf("error in opening file with args: %v, err: %v", os.Args[1], err)
	}

	words := 0
	inword := false

	b := bufio.NewReader(f)
	for {
		run, err := readByte(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("error in reading byte, err: %v", err)
		}

		if unicode.IsSpace(run) && inword {
			words++
			inword = false
		}

		inword = unicode.IsLetter(run)
	}

	log.Printf("num of words: %v", words)
}
