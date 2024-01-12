package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct{}

func (f *File) Read() {

}

func ReadFile(reader Reader) {
	// Reader 인터페이스 변수를 Closer 인터페이스로 타입 변환
	c := reader.(Closer)
	c.Close()
}

func main() {
	// [인터페이스 변환]
	// 2) 다른 인터페이스로 타입 변환하기
	//   - 구체화된 타입으로 변환할 때랑은 다르게, 변경되는 인터페이스가 변경 전 인터페이스를 포함하지 않아도 됨
	//   - 하지만 인터페이스가 가리키고 있는 실제 인스턴스가 변환하고자 하는 다른 인터페이스를 포함해야함

	file := &File{}
	ReadFile(file) // Error : interface conversion: *main.File is not main.Closer: missing method Close
}
