package main

// 런 타임 에러가 발생하지 않도록 CheckAndRun()수정

type Stringer interface {
	String()
}

type Reader interface {
	Read()
}

func CheckAndRun(stringer Stringer) {
	// r := stringer.(Reader)
	if r, ok := stringer.(Reader); ok {
		r.Read()
	}
}

func main() {
	var s Stringer
	CheckAndRun(s)
}
