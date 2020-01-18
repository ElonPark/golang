package main

import "fmt"

func main() {
	rawLiteral := `헬로우\n
월드\n
  안녕`
	fmt.Println(rawLiteral)

	interLiteral := "헬로우\n월드"
	fmt.Println(interLiteral)
}
