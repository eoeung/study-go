package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	// fmt.Printf("Age: %d\n", stringer.Age) // Stringer 인터페이스 변수로는 Age값에 접근할 수 없음
	s := stringer.(*Student) // *Student 타입으로 타입 변환
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	// [인터페이스 변환]
	// 1) 구체화된 다른 타입으로 변환
	//   - 인터페이스를 본래의 구체화된 타입으로 복원할 때 사용
	s := &Student{15}
	PrintAge(s)
}
