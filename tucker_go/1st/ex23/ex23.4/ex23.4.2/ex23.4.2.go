package main

// panic() 함수 선언
// func panic(interface{})

func main() {
	// [패닉 생성]
	// - panic() 함수의 인자는 interface{}이므로 모든 타입을 사용할 수 있음

	panic(42)
	panic("unreachable")
	// panic(fmt.Errorf("This is error num : %d", num))
	// panic(SomeType{SomeData})
}
