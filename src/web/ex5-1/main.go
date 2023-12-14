package main

import (
	"net/http"
	"web/ex5-1/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
