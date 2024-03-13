package benchmarkExperiment

import (
	"log"
	"math/rand"
	"testing"
)

func BenchmarkFib20(b *testing.B) {
	//var r int
	for n := 0; n < b.N; n++ {
		Fib(20)
	}
	//Result = uint64(r)
}

//
//func BenchmarkFib28(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		Fib(28)
//	}
//}
//
//func genStringNChars(n int) string {
//	var str string
//	for i := 0; i < n; i++ {
//		str += "a"
//	}
//
//	return str
//}
//
//func BenchmarkExpensive(b *testing.B) {
//	genStringNChars(100000)
//	b.ResetTimer()
//	for n := 0; n < b.N; n++ {
//		Fib(20)
//	}
//}
//
//func BenchmarkAllocs(b *testing.B) {
//	b.ReportAllocs()
//	genStringNChars(100000)
//	for n := 0; n < b.N; n++ {
//		Fib(20)
//	}
//}
//
//func TestTrue(t *testing.T) {
//	assert.Equal(t, true, true)
//}

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

var Result uint64

func BenchmarkPopcnt(b *testing.B) {
	//now := time.Now()
	var r uint64
	for i := 0; i < b.N; i++ {
		r = popcnt(rand.Uint64())
	}
	Result = r
	//log.Printf("benchmark time: %v", time.Since(now))
}

func BenchmarkFibWrong(b *testing.B) {
	log.Printf("b.N: %v", b.N)
	Fib(b.N)
}

func BenchmarkFibWrong2(b *testing.B) {
	log.Printf("b.N: %v", b.N)
	for i := 0; i < b.N; i++ {
		log.Printf("%v", i)
		Fib(i)
	}
}
