package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Foo")
}

func main() {
	// 핸들러 등록
	// Path에 대한 Request가 왔을 때, 함수를 실행
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Bar")
	})

	http.Handle("/foo", &fooHandler{})

	// PORT, ServeMux
	http.ListenAndServe(":3000", nil)
}
