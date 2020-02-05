package main

type dict struct {
	data map[int]string
}


// 구조체의 필드가 사용 전에 초기화되여하 하는 경우, 생성자 함수를 사용할 수 있다.
// 애를 들어 구조체의 필드가 map 타입인 경우 map을 사전에 미리 초기화 해놓으면,
// 외부 구조체 사용자가 매번 map을 초기화 해야한다는 것을 기억할 필요가 없다.
// 생성자 함수는 struct를 리턴하는 함수로서 그 함수 본문에서 필요한 필드를 초기화한다.

// 생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d //포인터 전달
}

func main() {
	// 직접 객체를 생성하지 않고 newDict()을 호출하여
	dic := newDict() // 생성자 호출
	// 이미 초기화된 data 맵 필드를 사용한다.
	dic.data[1] = "A"
}
