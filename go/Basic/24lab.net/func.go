package main

import "fmt"

func main() {
	msg := "Hello"
	say1(&msg)
	say(msg)
	say(msg)
	say2("This", "is", "a", "book")

	count, total := sum(1, 7, 3, 5, 9)
	fmt.Println(count, total)
	count2, total2 := sum2(1, 7, 3, 5, 9)
	fmt.Println(count2, total2)
}

func say(msg string) {
	fmt.Println(msg)
	msg = "end"
}

func say1(msg *string)  {
	fmt.Println(*msg)
	*msg = "Changed"
}

func say2(msg ...string)  {
	for _, s := range msg {
		fmt.Println(s)
	}
}

func sum(nums ...int) (int, int) {
	s := 0      // 합계
	count := 0  // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sum2(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}