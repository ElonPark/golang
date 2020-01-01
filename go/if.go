package main

import "fmt"

func main() {
	k := 10
	if k == 1 {
		fmt.Println("One")
	} else if k == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Other")
	}

	max := 10
	i := 3
	if val := i * 2; val < max {
		println(val)
	}
}
