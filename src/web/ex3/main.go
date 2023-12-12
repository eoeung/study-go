package main

import (
	"goproject/ex3/myapp" // 실행하면 잘 되는데, 에러 발생 → ???
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
