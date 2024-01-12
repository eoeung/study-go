package main

import "fmt"

type Stringer interface { // Stringer 인터페이스 선언
	String() string
	// 매개변수 없이 string타입을 반환하는 String()메서드를 포함한 모든 타입은, Stringer 인터페이스로 사용될 수 있음
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	// [인터페이스]
	// - 구현을 포함하지 않은 메서드 집합
	// - 인터페이스 구현 여부를 덕 타이핑으로 확인
	//   → 덕 타이핑 : 그 타입이 인터페이스에 해당하는 메서드를 가지고 있는지 판단

	// - 추상화된 객체로 상호작용 할 수 있음
	//   → 메서드 구현을 포함한 구체화된 객체가 아님

	// [인터페이스 작성 시 유의 사항]
	// 1) 메서드는 반드시 메서드명이 있어야 함
	// 2) 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없음 → 오버로딩 X
	// 3) 인터페이스에서는 메서드 구현을 포함하지 않음

	// Ex) 잘못된 예
	// type Sample interface {
	// 	String() string    // duplicate method String
	// 	String(int) string // duplicate method String
	// 	_(x int) // methods must have a unique non-blank name
	// }

	student := Student{"철수", 12}
	var stringer Stringer = student // Student 타입 student를 가지고 있음
	// Student 타입은 String() 메서드를 포함한다.
	// → Student 타입은 Stringer 인터페이스로 사용될 수 있음

	fmt.Printf("%s ::: %T ::: %T\n", stringer.String(), stringer, student)
}
