package main

import "fmt"

// http://golang.site/go/article/18-Go-인터페이스

// 빈 인터페이스
// Empty interface는 메소드를 전혀 갖지 않는 빈 인터페이스이다.
// Go의 모든 Type은 적어도 0개의 메소드를 구현하므로,
// 흔히 Go에서 모든 Type을 나타개 위한 빈 이터페이스를 사용한다.
// 다른 언터에서 Dynamic Type이라 볼수 있다.
// 마치 Java에서 최상위 부모 클래스인 Object처럼

func main() {
	var x interface{}
	x = 1
	x = "Tom"

	printIt(x)


	// Interface type의 x와 타입 T에 대하여 x.(T)로 표현했을 때,
	// 이는 x가 nil이 아니며, x는 T 타입에 속한다는 점을 확인(assert)하는 것으로
	// 이러한 표현을 "Type Assertion"이라 부른다.
	// 만약 x가 nil 이거나 x의 타입이 T가 아니라면, 런타임 에러가 발생할 것이고,
	// x가 T 타입인 경우는 T 타입의 x를 리턴한다.
	// 즉, 아래 예제에서 변수 j는 a.(int)로부터 int형 변수 j가 된다.
	var a interface{} = 1
	i := a       // a와 i 는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1

	printIt(i)
	printIt(j)
}

func printIt(v interface{})  {
	fmt.Println(v)
}