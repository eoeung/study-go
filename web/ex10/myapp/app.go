package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET '/' ::: Hello World")
	fmt.Fprint(w, "Hello World")
}

func tttHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET '/ttt' ::: tttttttttttt")
	fmt.Fprint(w, "tttttttttttt")
}

func NewHandler() http.Handler {
	// net/http로 구현한 기본적인 mux
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", indexHandler)
	// return mux

	// gorilla mux 구현
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/ttt", tttHandler)
	return mux
}
