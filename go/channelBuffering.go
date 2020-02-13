package main

import "fmt"

func main() {
	/*
		c := make(chan int)
		c <- 1   //수신루틴이 없으므로 데드락
		fmt.Println(<-c) //코멘트해도 데드락 (별도의 Go루틴이 없기 때문에)

		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.main()
		        /go/channelBuffering.go:7 +0x59
	*/

	ch := make(chan int, 1) // 1개의 정수형을 갖는 버퍼채널

	//수신자가 없더라도 보낼 수 있다.
	ch <- 101

	fmt.Println(<-ch)
}
