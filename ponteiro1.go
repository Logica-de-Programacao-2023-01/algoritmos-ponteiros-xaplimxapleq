package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := 1
	n := 4
	var ptr *int = &x
	var ptr2 *int = &n
	soma := nnumeros(ptr, ptr2)
	fmt.Println(soma)
}

func nnumeros(ptr *int, ptr2 *int) int {
	soma := 0
	currentPtr := ptr

	for {
		soma += *currentPtr

		if currentPtr == ptr2 {
			break
		}

		if currentPtr < ptr2 {
			currentPtr = incrementPtr(currentPtr)
		} else {
			currentPtr = decrementPtr(currentPtr)
		}
	}

	return soma
}

func incrementPtr(ptr *int) *int {
	address := uintptr(unsafe.Pointer(ptr))
	address += unsafe.Sizeof(*ptr)
	return (*int)(unsafe.Pointer(address))
}

func decrementPtr(ptr *int) *int {
	address := uintptr(unsafe.Pointer(ptr))
	address -= unsafe.Sizeof(*ptr)
	return (*int)(unsafe.Pointer(address))
}
