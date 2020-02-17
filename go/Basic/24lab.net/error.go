package main

import (
	"fmt"
	"log"
	"os"
)

// 이 인터페이스를 구현하는 커스텀 에러 타입을 만들 수 있다.
type MyError interface {
	Error() string
}

func main() {
	f, errr := os.Open("${HOME}/Downloads/test.txt")
	if errr != nil {
		log.Fatal(errr.Error())
	}
	fmt.Println(f.Name())


	//otherFunc()를 호출한 후 error가 err로 리턴되었을 때,
	//이 err의 타입별로 다른 처리를 하는 것을 볼 수 있다
	//(주: switch 문에서 "변수명.(type)"의 방식으로 타입 체크를 한다).
	//디폴트의 경우는 에러타입이 nil인 경우로서 에러가 없는 경우이고,
	//에러가 있으면 다음 case문에서 그 에러타입이 MyError인지를 체크하고,
	//아니면 다음 case에서 일반 에러 케이스를 처리한다.
	//모든 에러는 error 인터페이스를 구현하므로 마지막 case문은 모든 에러에 적용된다.

	//_, err := otherFunc()
	//switch err.(type) {
	//default: // 에러 없음
	//	fmt.Println("OK")
	//case MyError:
	//	log.Print("Log my error")
	//case error:
	//	log.Fatal(err.Error())
	//}
}


