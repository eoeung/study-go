// 진짜 이해하고 있는지 머릿속으로 생각하면서 코딩해보기
package main

import (
	"fmt"
	"net/http"
)

// 함수형 등록 (함수 만들어서 사용)
func temp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Func")
}

// 구조체형 등록
type car struct {
	name  string
	color string
}

// car에 대한 메서드 생성
func (c *car) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "carName : %s, carColor : %s", c.name, c.color)
}

func main() {
	// ###################### [완전 기본적인 방법] ######################
	// 함수형 등록
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World")
	// })

	// 함수형 등록 (함수 만들어서 사용)
	// http.HandleFunc("/func", temp)

	// 구조체형 등록
	// http.Handle()은 string, Handler를 인자로 넘거야한다.
	// tc := car{name: "BMW", color: "검정"} // 구조체 초기화
	// aaa := &tc                          // 구조체 포인터를 aaa에 저장
	// http.Handle("/car", aaa)
	// http.Handle("/car", &car{}) // &car{} : 새로운 car 구조체의 포인터를 생성 → 포인터 생성

	// 서버 설정 (mux 미사용)
	// http.ListenAndServe(":3000", nil)

	// ###################### [최대한 클린코드 및 mux 사용] ######################
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/func", temp)
	mux.Handle("/car", &car{name: "포르쉐", color: "빨강"}) // 임의로 구조체 필드에 값 넣어서 진행

	// 서버 설정 (mux 사용)
	http.ListenAndServe(":3000", mux)
}
