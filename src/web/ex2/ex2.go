package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 구조체 → 인스턴스를 이용해 등록 (Handle)
type fooHandler struct{}

// User 구조체 생성
type User struct {
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) // request body의 값을 읽어서, Decode 하겠다는 의미

	// 에러 처리
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request : ", err)
		return
	}
	user.CreatedAt = time.Now() // 현재 시간

	fmt.Println(time.Now()) // Ex. 2023-12-11 18:01:36.6789528 +0900 KST m=+16.372123201

	data, _ := json.Marshal(user)                      // json 형태로 변환
	w.Header().Add("content-type", "application/json") // json 형태로 응답한다고 헤더에 선언
	w.WriteHeader(http.StatusOK)                       // 200 상태를 헤더에 선언
	fmt.Fprint(w, string(data))
}

// 함수로 등록 (HandlerFunc)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // QueryString
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func main() {
	// mux : 라우터
	mux := http.NewServeMux()
	// 핸들러 등록
	// Path에 대한 Request가 왔을 때, 함수를 실행
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})

	// PORT, ServeMux
	http.ListenAndServe(":3000", mux) // mux(라우터)를 등록
}
