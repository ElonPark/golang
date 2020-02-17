package main

import "fmt"

func main() {
	var k0 int = 10
	var p = &k0
	fmt.Println(p, "=", *p)
}
