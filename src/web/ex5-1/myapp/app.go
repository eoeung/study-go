// 핸들러를 만들어주는 패키지
package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

//	func getUserInfo89Handler(w http.ResponseWriter, r *http.Request) {
//		fmt.Println(r.URL.Path) // /users/89
//		fmt.Fprint(w, "User Id:89")
//	}
func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path) // /users/89
	vars := mux.Vars(r)
	fmt.Fprint(w, "User Id:", vars["id"])
}

// myapp 핸들러 생성
func NewHandler() http.Handler {
	mux := mux.NewRouter() // gorilla mux 사용
	// mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	// mux.HandleFunc("/users/89", getUserInfo89Handler) // 하드코딩
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)
	return mux
}
