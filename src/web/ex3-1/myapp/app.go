package myapp

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
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Email     string    `json:"Email"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	// 에러 처리
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request : ", err)
		return
	}
	user.CreatedAt = time.Now()

	fmt.Println(time.Now())

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json") // json 형태로 응답한다고 헤더에 선언
	w.WriteHeader(http.StatusOK)                       // 200 상태를 헤더에 선언
	fmt.Fprint(w, string(data))
}

// 함수로 등록 (HandlerFunc)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello World") // 테스트할 때, "Hello World\n"로 들어옴
	fmt.Fprint(w, "Hello World")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // QueryString
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})
	return mux
}

func FuncA(param string) {
	//......
}
