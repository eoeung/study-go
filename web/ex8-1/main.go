package main

import (
	"net/http"
	"web/ex8-1/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
