package main

import (
	"net/http"
	"study-go/web/ex7-1/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
