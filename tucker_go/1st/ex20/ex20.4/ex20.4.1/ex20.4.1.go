package main

// 1) 포함된 인터페이스
// Ex) Reader와 Writer 인터페이스를 선언하고, 이 둘을 포함한 ReadWriter 인터페이스를 선언하는 예제
type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

// Reader와 Writer의 인터페이스 메서드 집합을 모두 포함한 ReadWriter 인터페이스
type ReadWriter interface {
	Reader // Reader의 메서드 집합 포함
	Writer // Writer의 메서드 집합 포함
}

func main() {
	// [인터페이스의 3가지 다른 기능]
	// 1) 포함된 인터페이스
	//   - 인터페이스를 포함하는 인터페이스
	//   - 구조체에서도 다른 구조체를 포함된 필드로 가질 수 있는 것과 똑같다.
}
