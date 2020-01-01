package main

import (
	"fmt"
)

func main() {
	var a int
	var f float32
	a = 10
	f = 12.0
	fmt.Println(a, f)

	var i, j, k int = 1, 2, 3
	var i0 = 1
	var s = "Hi"
	fmt.Println(j, k)
	fmt.Println(i, i0, s)

	const c int = 10
	const s0 string = "Hi"
	const c0 = 10
	const c1 = "H1"

	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)
	fmt.Println(Amex)

	const (
		Apple = iota
		Grape
		Orange
	)
	fmt.Println(Apple, Orange)
}
