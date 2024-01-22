package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	// [별칭 리시버 타입]
	// - 모든 로컬 타입이 리시버 타입으로 가능
	//   → int 같은 내장 타입도 별칭 타입을 활용해서 메서드를 가질 수 있음

	var a myInt = 10       // myInt 타입 변수
	fmt.Println(a.add(30)) // myInt 타입의 add() 메서드 호출

	var b int = 20
	fmt.Println(myInt(b).add(50)) // int타입을 myInt타입으로 변환
	// 1) myInt 타입은 int의 별칭 타입이다.
	// 2) 하지만, 엄연히 다른 타입이기 때문에 int 타입은 myInt의 add()메서드를 사용할 수 없다.
	// 3) myInt 타입은 int 타입이므로, myInt()와 같이 별칭 타입 간 타입 변환을 지원하는 것을 이용했다.

	// var c string = "30"
	// fmt.Println(myInt(c).add(20)) // cannot convert c (variable of type string) to type myInt
	// string 타입을 바로 myInt 타입으로 변환하는 것은 에러가 발생했다.
}
