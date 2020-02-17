package main

import "fmt"

func main() {
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)

	add := func(i int, j int) int {
		return i + j
	}

	r1 := calc(add, 10, 20)
	fmt.Println(r1)

	//파라미터/에 직접 익명함수를 정의
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	fmt.Println(r2)

	r3 := calc2(
		func(x int, y int) int {
			return x - y
		},
		20, 10)
	fmt.Println(r3)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

type calculator func(int, int) int

func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}
