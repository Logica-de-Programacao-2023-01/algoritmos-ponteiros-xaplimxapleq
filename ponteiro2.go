package main

import "fmt"

func main() {
	a := 3
	var ptr *int = &a
	parimpar(ptr)
	fmt.Println(a)
}

func parimpar(ptr *int) {
	if *ptr%2 != 0 {
		*ptr = 1
	} else {
		*ptr = 0
	}
}
