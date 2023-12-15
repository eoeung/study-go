package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func NewHandler() http.Handler {
	// net/http로 구현한 기본적인 mux
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", indexHandler)
	// return mux

	// gorilla mux 구현
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
	return mux
}
