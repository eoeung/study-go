package main

import (
	"log"
	"net/http"
	"web/ex21/app"
)

func main() {
	// configuration 관련 코드는 최대한, 패키지가 아닌 최상단에 위치하는 것이 좋음
	m := app.MakeHandler("./test.db")
	defer m.Close()

	log.Println("Started")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
