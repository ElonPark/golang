package main

import (
	"fmt"
	"time"
)

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT: // 레이블로 이동 한후 자신이 빠져나온 루프 다음 문장을 실행하게 된다.
	for {
		/*
			Go의 select문은 복수 채널들을 기다리면서 준비된 (데이터를 보낸) 채널을
			실행하는 기능을 제공한다.

			select문은 여러 개의 case문에서 각각 다른 채널을 기다리다가
			준비가 된 채널 case를 실행한다. 만약 복수의 채널에 신호가 오면
			Go 런타임이 랜덤하게 그 중 한 개를 선택한다. 하지만 default문이 있으면,
		case문 채널이 준비되지 않더라도 계속 대기하지 않고 바로 default문을 실행한다.
		*/
		select {
		case <-done1:
			fmt.Println("run1 완료")
		case <-done2:
			fmt.Println("run2 완료")
			break EXIT // 이 문장으로 인해 for 루프를 빠져나와 EXIT 레이블로 이동한다.
		}
	}
	//이 경우는 main() 함수의 끝에 다다르게 됨
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
