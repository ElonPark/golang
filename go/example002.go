package main

import "fmt"

func fac(n int) int {
	if n <= 0 {
		return 1
	}

	return n * fac(n-1)
}

func facItr(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}

	return result
}

func main() {
	fmt.Println(fac(5))
	fmt.Println(facItr(5))
}
