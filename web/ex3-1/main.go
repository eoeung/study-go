package main

import (
	"net/http"
	"web/ex3-1/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
