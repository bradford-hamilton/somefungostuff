package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/bradford-hamilton/somefungostuff/adder"
)

func main() {
	// the adder go package defines the interface, while
	// add_amd64.s implements the function in assembly
	num := adder.Add(1, 2)
	fmt.Printf("num: %d\n", num)

	// take byte slice and turn it into a string using custom func
	str := BytesToString([]byte{'b', 'r', 'a', 'd'})
	fmt.Println("string:", str)
}

// BytesToString turns a []byte into a string with 0 MemAllocs and 0 MemBytes.
func BytesToString(bytes []byte) (s string) {
	if len(bytes) == 0 {
		return ""
	}
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	str.Data = uintptr(unsafe.Pointer(&bytes[0]))
	str.Len = len(bytes)
	return
}
