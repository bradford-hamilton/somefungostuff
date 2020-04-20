package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// benchmarkConvertByteSliceToString - 27.8 ns/op
// MemAllocs: 43246367
// MemBytes: 2767766760

// benchmarkBytesToStringCustomFunc - 1.65 ns/op
// MemAllocs: 0
// MemBytes: 0

func main() {
	res := testing.Benchmark(benchmarkConvertByteSliceToString)
	fmt.Printf("%s\n%#[1]v\n", res)
	fmt.Printf("MemAllocs: %d\n", res.MemAllocs)
	fmt.Printf("MemBytes: %d\n", res.MemBytes)

	res = testing.Benchmark(benchmarkBytesToStringCustomFunc)
	fmt.Printf("%s\n%#[1]v\n", res)
	fmt.Printf("MemAllocs: %d\n", res.MemAllocs)
	fmt.Printf("MemBytes: %d\n", res.MemBytes)
}

var result string
var bytes = []byte(`just a buncha bytes hangin out, nothing to see here.`)

func benchmarkConvertByteSliceToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := bytesToStringWithString(bytes)
		result = s
	}
}

func bytesToStringWithString(bytes []byte) string {
	return string(bytes)
}

func benchmarkBytesToStringCustomFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := customBytesToString(bytes)
		result = s
	}
}

func customBytesToString(bytes []byte) (s string) {
	if len(bytes) == 0 {
		return ""
	}
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	str.Data = uintptr(unsafe.Pointer(&bytes[0]))
	str.Len = len(bytes)
	return
}
