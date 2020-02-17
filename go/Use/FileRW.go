package main

import (
	"io"
	"os"
)

func main() {
	// 입력파일 열기
	fi, err := os.Open("/Users/elon/Develop/temp/1.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	// 출력파일 생성
	fo, err := os.Create("/Users/elon/Develop/temp/2.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 2014)

	// 루프
	for {
		// 읽기
		cnt, err := fi.Read(buff)
		
		if err != nil && err != io.EOF {
			panic(err)
		}

		// 끝이면
		if cnt == 0 {
			break
		}

		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}
