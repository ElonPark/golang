package main

import "fmt"

// Rect - struct 정의
type Rect struct {
	width, height int
}

// Rect의 area() 메소드
// 밸류 리시버
// 밸류 리시버는 데이터를 복사하여 전달하며 메소드 내에서 필드가 값이 변경되어도
// 호출자의 데이터는 변경되지 않는다.
func (receiver Rect) area() int {
	receiver.width++ //변경 내용은 호출자에 반영되지 않는다.
	return receiver.width * receiver.height
}

// 포인터 리시버
// 메소드 내의 필드 값 변경이 호출자에 반영된다.
func (receiver *Rect) area2() int {
	receiver.width++ // 변경 내용이 호출자에 반영된다.
	return receiver.width * receiver.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() // 메소드 호출
	fmt.Println(rect, area)

	area = rect.area2() // width 값 변경 됨.
	fmt.Println(rect, area)
}