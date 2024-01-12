package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct{}

func (s *Student) String() string {
	return "Student"
}

type Actor struct{}

func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	// [런타임 에러가 발생하는 경우]
	// - 변환하려는 타입이 인터페이스를 포함하고 있는 상황
	//   - 실제 인터페이스 변수가 가리키는 인스턴스가 변환하려는 타입이 아닌 경우
	actor := &Actor{}
	ConvertType(actor) // panic: interface conversion: main.Stringer is *main.Actor, not *main.Student
	// 문법적인 오류는 없기 때문에 컴파일 에러는 발생하지 않음

	// 1) ConvertType() 함수 인수인 stringer 인터페이스 변수는 *Actor 타입 인스턴스를 가리키고 있음
	// 2) *Student 타입 변환을 시도하면 런타임 에러 발생
}
