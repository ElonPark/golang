package main

import "fmt"

func main() {
	//정수형 3개 요소를 갖는 배
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a[1])

	a2 := [3]int{4, 5, 6} // 초기값과 같이 배열 선언
	a3 := [...]int{9, 8, 7} // 배열 크기 자동으로
	fmt.Println(a2)
	fmt.Println(a3)

	// 다차원 배열
	var multiArray [3][4][5] int
	multiArray[0][1][2] = 10
	fmt.Println(multiArray)

	var a4 = [2][3]int {
		{1,2,3},
		{4,5,6}, // 마지막 원소는 콤마가 없으면 안됨.
		// 스위프트는 있던지 말던지 둘다 상관 없는데..
	}
	fmt.Println(a4[1][2])
}
