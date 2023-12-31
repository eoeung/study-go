package main

import (
	"log"
	"net/http"
	"study-go/web/ex17/app"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
