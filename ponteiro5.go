package main

import (
	"fmt"
	"math"
)

func mediaPi(n float64) float64 {
	return (math.Pi + n) / 2
}

func main() {
	x := 15.16
	result := mediaPi(x)
	fmt.Println(result)
}
