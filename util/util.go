package util

import (
	"fmt"
	"reflect"
	"unsafe"
)

func PrintAsBinary(a any) {
	type iface struct {
		t, v unsafe.Pointer
	}
	p := uintptr((*(*iface)(unsafe.Pointer(&a))).v)

	t := reflect.TypeOf(a)

	for i := 0; i < int(t.Size()); i++ {
		n := *(*byte)(unsafe.Pointer(p))
		fmt.Printf("%08b ", n)
		p += unsafe.Sizeof(n)
	}

	fmt.Print("\n")
}
