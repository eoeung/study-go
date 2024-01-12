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
	// 타입 변환 가능 여부 확인
	if c, ok := reader.(Closer); ok {
		c.Close()
	}

	// c, ok := reader.(Closer)
	// if ok {
	// 	c.Close()
	// }
}

func main() {
	// [타입 변환 가능 여부를 확인]
	// - 런타임 에러가 발생하지 않도록 타입 변환 가능 여부를 확인한다.

	// var a interface
	// t, ok := a.(ConcreteType)
	// t: 타입 변환한 결과 → 만약 타입 변환에 실패할 경우, ConcreteType의 기본값이 온다.
	// ok: 타입 변환 성공 여부

	file := &File{}
	ReadFile(file)
}
