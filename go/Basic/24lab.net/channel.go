package main

import (
	"fmt"
	"time"
)

func main() {
	// 정수형 채널 생성한다.
	ch := make(chan int)

	go func() {
		ch <- 123 // 채널에 123을 보냄
	}()

	var i int
	i = <- ch // 채널로부터 123을 받는다.
	fmt.Println(i)

	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			if i == 5 {
				time.Sleep(time.Second * 2)
			}
			fmt.Println(i)
		}
		done <- true
	}()

	// 위의 Go루틴이 끝날 때까지 대기
	<-done
}