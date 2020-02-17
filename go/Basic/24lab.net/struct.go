package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	p := person{}

	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	p1 := person{"Bob",  20}
	fmt.Println(p1)

	// 필드명을 지정하여 초기화하는 경우, 만약 일부 필드가 생량될 경우
	// 생략된 필드들은 Zero Value(int = 0, float = 0.0, string = "", 포인터 = nil)
	// 를 가진다.
	p2 := person{name: "Sean", age: 50}
	fmt.Println(p2)

	// new()를 사용하면 모든 필드를 Zero value로 초기화하고
	// person 객체의 포인터(*person) 리턴한다.
	p3 := new(person)
	p3.name = "Lee" // p3가 포인터라도 . 을 사용한다.
	//Go에서 struct를 mutable 개체로서 필드 값이 변화될 경우 별로도 새 개체를 만들지
	//않고 해당 개체 메모리에서 직접 변경된다. 하지만 개체를 다른 함수의 파라터로
	//남간다면, Pass by Value에 따라 객체를 복사하여 전달한다.
	//만약 Pass by Reference로 전달하려면 struct의 포인터를 전달해야한다.
}
