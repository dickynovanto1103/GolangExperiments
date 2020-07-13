package main

import (
	"strings"
	"testing"
)

func genString(len int) string {
	s := make([]byte, len)
	for i:=0;i<len;i++ {
		s[i] = byte('a')
	}
	return string(s)
}

var s1 string = genString(1000000)
var s2 string = genString(1000000)

func Benchmark_EqualFold(b *testing.B) {
	s1 := genString(1000000)
	s2 := genString(1000000)
	strings.EqualFold(s1, s2)
}

func Benchmark_ToUpper(b *testing.B) {
	_ = strings.ToUpper(s1) == strings.ToUpper(s2)

}