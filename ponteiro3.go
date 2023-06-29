package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func main() {
	s := "Oi"
	inverted := stringutil.Reverse(s)
	fmt.Println(inverted)
}
