package main

import "fmt"

func isPrimo(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primos(s *[]int, n int) {
	if n <= 0 {
		return
	}

	i := 2
	for {
		if isPrimo(i) {
			*s = append(*s, i)
			n--
		}

		if n == 0 {
			break
		}

		i++
	}
}

func main() {
	conj := []int{}
	primos(&conj, 10)
	fmt.Println(conj)
}
