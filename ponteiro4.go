package main

import (
	"fmt"
)

func somaDoisUltimosAlgarismos(n *int) {
	ultimoAlgarismo := *n % 10
	*n /= 10
	penultimoAlgarismo := *n % 10

	soma := ultimoAlgarismo + penultimoAlgarismo

	*n = soma
}

func main() {
	x := 28654

	somaDoisUltimosAlgarismos(&x)

	fmt.Println(x)
}
