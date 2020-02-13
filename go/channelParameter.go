package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)

	/*
		채널을 오픈한 후 데이터를 송신한 후,
		close() 함수를 사용하여 채널을 닫을 수 있다.
		채널을 닫게 되면, 해당 채널로는 더이상 송신을 할 수 없지만,
		채널을 닫힌 이후에도 계속 수신은 가능하다.

		채널에 수신에 사용하는  <-ch은 두개의 리턴 값을 갖는데,
		첫째는 채널 메시지이고, 두번째는 수신이 제대로 되었는가를 나타낸다.
		만약 채널이 닫혔다면, 두번째 값은 false를 리턴한다.
	*/
	ch2 := make(chan int, 2)
	// 채널에 송신
	ch2 <- 1
	ch2 <- 2

	// 채널을 닫는다.
	close(ch2)

	// 채널 수신
	//fmt.Println(<-ch2)
	//fmt.Println(<-ch2)

	/*if _, success := <- ch2; !success {
		fmt.Println("더이상 데이터 없음.")
	}*/

	/*
		수신자는 채널이 닫히는 것을 체크하며 계속 루프를 돌게 된다.
		채널 range문은 range 키워드 다음의 채널로부터 계속 수신하다가 채널이 닫힌 것을
		감지하면 for 루프를 종료한다.
	*/
	// 방법1
	/*for {
		if i, success := <-ch; success {
			fmt.Println(i)
		} else {
			break
		}
	}*/

	// 방법2
	// 방법1의 표현과 동일한 채널 range 문
	for i := range ch2 {
		fmt.Println(i)
	}
}

// 송신
func sendChan(ch chan<- string) {
	ch <- "Data"
	//x := <- ch //에러 발생. 송신 정용 채널에서 수신을 시도하므로 에러가 발생한다.
	/*
		# command-line-arguments
		./channelParameter.go:13:7:
		invalid operation: <-ch (receive from send-only type chan<- string)
	*/
}

// 수신
func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
