package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

// func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo")
}

func main() {
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	// Handler를 함수 형태로 직접 등록할 때 사용
	// URL 패턴 인자인 pattern에 대해서, 요청에 대한 핸들러 함수리터럴를 등록
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ResponseWriter는 클라이언트에게 응답을 작성하는데 사용됨
		// http.Request는 클라이언트의 요청 정보를 나타냄

		// ※ *http.Request로 사용한 이유
		// struct(구조체) 타입은 구조체의 인스턴스가 함수에 전달될 때 마다 복사가 발생한다.
		// 포인터를 사용해서 구조체의 주소를 전달하여, 함수 내에서 구조체를 직접 조작할 수 있게 한다.
		// 메모리 사용도 최적화 할 수 있음
		// 만약 포인터를 사용하지 않는다면,
		// 복사된 인스턴스가 함수 내에서만 변경되고 실제 전달한 인스턴스 값은 그대로임

		fmt.Fprint(w, "Hello World")
		// fmt.Fprint() 함수는 형식화된 문자열을 특정한 io.Writer에 쓰기 위한 함수
		// func Fprint(w io.Writer, a ...interface{}) (n int, err error)

		// fmt.Print() 함수는 형식화된 문자열을 표준 출력에 쓰기 위한 함수
		// func Print(a ...interface{}) (n int, err error)

		// HTTP 응답을 작성해야하므로 fmt.Fprint()를 사용
		// fmt.Print()를 사용하면 표준 출력에 문자열이 출력되고, 클라이언트에게 전달되지 않음

		// 실제 fmt.Print()사용하면, 브라우저는 흰색으로 아무것도 출력되지 않고,
		// 아래와 같이 서버쪽 console에 출력됨
		/*
			&{0xc00016c000 0xc0000a0000 {} 0xb42b20 false false false false {{} 0} {0 0}
			0xc00009a080 {0xc0000c2000 map[] false false} map[] false 0 -1 0 false false false
			[] {{} 0} [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
			[0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc0000ba070 {{} 0}}
			Hello World&{0xc00016c000 0xc0000a0100 {} 0xb42b20 false false false false
			{{} 0} {0 0} 0xc00009a080 {0xc0000c21c0 map[] false false} map[] false 0 -1 0
			false false false [] {{} 0} [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
			[0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc0000ba150 {{} 0}}Hello World
		*/
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	// Handler, 인스턴스 형태로 등록할 때는 http.Handle을 사용
	// 인터페이스를 http.ServeHTTP 함수를 사용해서 만들어서 사용
	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil)
	// func ListenAndServe(addr string, handler Handler) error
	// 첫 번째 인자에 해당하는 PORT에서 Request를 Listen할 것을 지정
	// 두 번째 인자는 어떤 ServeMux를 사용할 지 지정
	//   → http.Handler 인터페이스를 따르는 객체
}
