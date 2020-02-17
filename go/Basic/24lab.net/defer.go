package main

import (
	"fmt"
	"os"
)

// Swift defer와 비슷함.

func main() {
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	fmt.Println(len(bytes))
}
