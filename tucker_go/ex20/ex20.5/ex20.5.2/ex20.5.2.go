package main

type Stringer interface {
	String() string
}

type Student struct{}

func main() {
	// [컴파일 에러가 발생하는 경우]
	// - 인터페이스를 구현하지 않은 구조체로 타입 변환 시도하는 경우

	var stringer Stringer
	stringer.(*Student)
	/*
		.\ex20.5.2.go:11:2: impossible type assertion: stringer.(*Student)
		*Student does not implement Stringer (missing method String)
	*/
}
