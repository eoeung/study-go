package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	log.Printf("postMessageHandler %v %v", name, msg)
}

func main() {
	// gorilla/pat 사용
	mux := pat.New()
	mux.Post("/messages", postMessageHandler)

	// negroni 사용
	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
