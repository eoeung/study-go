package main

import "fmt"

// 2) 빈 인터페이스 interface{}를 인수로 받기
func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("No Supported type : %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	// [인터페이스의 3가지 다른 기능]
	// 2) 빈 인터페이스 interface{}를 인수로 받기
	//   - 메서드를 가지고 있지 않은 빈 인터페이스
	//   - 가지고 있어야 할 메서드가 하나도 없음
	//     → 어떤 값이든 받을 수 있다.

	PrintVal(10)          // v is int 10
	PrintVal(3.14)        // v is float64 3.140000
	PrintVal("Hello")     // v is string Hello
	PrintVal(Student{15}) // No Supported type : main.Student:{15}

}
