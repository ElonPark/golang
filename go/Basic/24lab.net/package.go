package main

import "fmt"
import "24lab.net/testlib"

// $ cd ~/go/src/
// $ mkdir -p 24lab.net/testlib
// 커스텀 라이브러리용 파일 작성
// $ go install # 파일내에서 사용할 라이브러리 인스톨

func main() {
	song := testlib.GetMusic("Alicia Keys")
	fmt.Println(song)

	//http://golang.site/go/article/15-Go-패키지
}
