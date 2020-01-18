package main

func main() {
	forLoop()
	forLoopLikeAwhile()
	//infinityLoop()
	forRange()
	ifAndGoto()
	breakLabel()
}

func forLoop() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)
}

func forLoopLikeAwhile() {
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)
}

func infinityLoop() {
	for {
		println("Infinity Loop")
	}
}

func forRange() {
	names := []string{"홍길동", "전우치", "이순신"}
	for index, name := range names {
		println(index, name)
	}
}

func ifAndGoto() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 1 {
		goto END
	}
END:
	println("END")
}

func breakLabel() {
	i := 0
L1:
	for {
		if i == 0 {
			break L1
		}
	}
	println("OK")
}
