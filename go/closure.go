package main

import "fmt"

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextValue()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	anotherNext := nextValue()
	fmt.Println(anotherNext())
	fmt.Println(anotherNext())
	fmt.Println(anotherNext())
}
