package main

import (
	"encoding/binary"
	"log"
)

func main() {
	tryLittleEndian()
	tryBigEndian()
	tryLittleEndianAppendUint32()
}

// learnt:
// 1. we can print value in hexadecimal using #
func tryLittleEndian() {
	buf := make([]byte, 4)
	x := 31
	log.Printf("%x %X %#x %#X", x, x, x, x)

	val := 16909060
	log.Printf("16909060 in hex: %#x", val)
	binary.LittleEndian.PutUint32(buf, uint32(val))
	log.Printf("%#v", buf)
}

func tryBigEndian() {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(16909060))
	log.Printf("%#v", buf)
}

func tryLittleEndianAppendUint32() {
	num := uint32(16909060) // 0x01020304 in hexadecimal
	var buf []byte
	buf = binary.LittleEndian.AppendUint32(buf, num)
	log.Printf("%#v\n", buf)
	newNum := binary.LittleEndian.Uint32(buf)
	log.Printf("%v", num == newNum)
}
