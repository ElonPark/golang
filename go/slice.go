package main

import "fmt"

func main() {
	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
	a[1] = 10
	fmt.Println(a)

	s := make([]int, 5, 10) // int 타입 슬라이스, 길이 5, 용량 10
	fmt.Println(s, len(s), cap(s))

	var s1 []int //nil 슬라이스
	if s1 == nil {
		fmt.Println("nil Slice")
	}
	fmt.Println(s1, len(s1), cap(s1))

	s2 := []int{-1, 0, 1, 2, 3, 4, 5, 6}
	fmt.Println(s2, len(s2), cap(s2))
	// 서브 슬라이스 - 파이썬과 동일하지만, 파이썬과 달리 음수 인덱스는 안됨
	s2 = s2[1:4] // 서브 슬라이스를 하게 되면 슬라이스 내부에 배열 포인터는
	// 인덱스가 1인 두번째 요소인 0을 가리키고, 길이는 3,
	// 용량은 배열의 두번째 요소부터 마지막 요소까지의 크기인 7을 갖게 된다.
	// 기존 배열의 두번째부터 마지막까지 [0, 1, 2, 3, 4, 5, 6] <- 7개
	fmt.Println(s2, len(s2), cap(s2))
	s2 = s2[2:] //용량은 기존 배열의 세번째부터 마지막까지 [2,3,4,5,6] <- 5개
	// 서브 슬라이스를 하여 변수에 할당하더라도 슬라이스 내부배열은 변하지 않는다.
	fmt.Println(s2, len(s2), cap(s2))

	s3 := []int{0, 1}
	//추가
	s3 = append(s3, 2)
	fmt.Println(s3)
	// 한번에 여러개 추가
	s3 = append(s3, 3, 4, 5, 6)
	fmt.Println(s3)

	/*
		append()가 슬라이스에 데이터를 추가할때 용량이 남아 있는 경우,
		용량 내에서 슬라이스 길이를 변경하여 데이터를 추가하고
		용량을 초과하는 경우, 현재 용량의 2배에 해당하는 새로운 배열을 생성하여
		기존 배열값등을 모두 새배열에 복제한 후 다시 슬라이스에 할당한다.
		* 스위프트 배열도 go의 슬라이스처럼 Array Doubling 방식을 사용해서
		용량이 초과되면 배열은 더 큰 메모리 영역 할당하고 추가할 요소를 새 스토리지에
		복사한다.
		새 스토리지 용량은 기존 배열 용량의 2배 <- go와 동일한 Array Doubling 방식
	*/
	sliceA := make([]int, 0, 3) //길이 0, 용량 3
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)

	// 슬라이스 병합
	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}

	// sliceC 뒤의 ...은 sliceC의 모든 요소를 붙인다는 것을 의미한다
	sliceB = append(sliceB, sliceC...)
	//슬라이스 C의 내부 배열 데이터들로 순서대로 치환됨. sliceC.. -> 4, 5, 6
	//sliceB = append(sliceB, 4, 5, 6) <- 이렇게 치환된다.

	fmt.Println(sliceB, len(sliceB), cap(sliceB))

	//슬라이스 복사
	/*
		슬라이스는 실제 배열을 가리키는 포인터 정보만 가지므스로
		복사를 하게 되면 source 슬라이 내부 배열의 데이터를
		target 슬라이스 내부 배열로 복제하는 것이다.
	*/
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source))
	copy(target, source)
	fmt.Println(target)
	fmt.Println(len(target), cap(target))

}
