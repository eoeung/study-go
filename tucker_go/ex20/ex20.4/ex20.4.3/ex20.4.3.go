package main

import "fmt"

// 3) 인터페이스 기본값 nil
type Attacker interface {
	Attack()
}

func main() {
	// [인터페이스의 3가지 다른 기능]
	// 3) 인터페이스 기본값 nil
	//   - 인터페이스 변수의 기본값 : nil
	//   - nil : 유효하지 않은 메모리 주소

	var att Attacker // 기본값은 nil
	// att.Attack()     // Error : panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("%T ::: %v\n", att, att) // <nil> ::: <nil>
	// → nil이 아닌지 확인해야함
	// ※ 인터페이스, 포인터, 함수 타입, 슬라이스, 맵, 채널, ...
}
