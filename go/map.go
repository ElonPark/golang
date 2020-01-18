package main

import "fmt"

func main() {
	// 선언
	var idMap map [int]string
	// 초기화
	idMap = make(map[int]string)
	fmt.Println(idMap)

	// 리터럴을 사용한 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB": "FaceBook",
	}
	// 키 체크
	val, exits := tickers["MSFT"]
	fmt.Println(val)
	if !exits {
		fmt.Println("No MSFT ticker")
	}
	fmt.Println(tickers)

	for key, value := range tickers {
		fmt.Println(key, value)
	}


	var m map[int]string
	m = make(map[int]string)
	// 추가 혹은 갱신
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	fmt.Println(str)

	/* 값이 없으면 레퍼턴스 타입인 경우는 nil
	밸류 타입은 zero 리턴, string = "", int = 0, bool = false */
	noData := m[999]
	fmt.Println(noData)
	fmt.Println("End")
	// 삭제
	delete(m, 777)

	intMap := make(map[int]int)
	fmt.Println(intMap[0])
	boolMap := make(map[int]bool)
	fmt.Println(boolMap[0])
}
