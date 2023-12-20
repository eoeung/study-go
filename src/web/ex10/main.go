package main

import (
	"log"
	"net/http"
	"time"
	"web/ex10/decoHandler"
	"web/ex10/myapp"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	// main 기능 : 서버 호출
	// sub 기능 : 로그 출력
	start := time.Now()
	log.Println("[LOGGER1] Started")                                           // 핸들러를 호출하기 전, start 로그 출력
	h.ServeHTTP(w, r)                                                          // Handler.ServeHTTP()를 구현
	log.Println("[LOGGER1] Completed time:", time.Since(start).Milliseconds()) // h.ServeHTTP(w, r)을 실행하고 종료까지 시간을 출력

}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	// main 기능 : 서버 호출
	// sub 기능 : 로그 출력
	start := time.Now()
	log.Println("[LOGGER2] Started")                                           // 핸들러를 호출하기 전, start 로그 출력
	h.ServeHTTP(w, r)                                                          // Handler.ServeHTTP()를 구현
	log.Println("[LOGGER2] Completed time:", time.Since(start).Milliseconds()) // h.ServeHTTP(w, r)을 실행하고 종료까지 시간을 출력
}

// func encrypt().. 등 부가 기능을 추가할 수 있음

func NewHandler() http.Handler {
	// h := myapp.NewHandler()
	// // NewDecoHandler 구현된 http handle (input http handle)
	// h = decoHandler.NewDecoHandler(h, logger)
	// h = decoHandler.NewDecoHandler(h, logger2)

	h1 := myapp.NewHandler()
	// NewDecoHandler 구현된 http handle (input http handle)
	h2 := decoHandler.NewDecoHandler(h1, logger)
	h3 := decoHandler.NewDecoHandler(h2, logger2)

	return h3

	/*
		2023/12/15 11:44:21 [LOGGER2] Started
		2023/12/15 11:44:21 [LOGGER1] Started
		2023/12/15 11:44:21 [LOGGER1] Completed time: 0
		2023/12/15 11:44:21 [LOGGER2] Completed time: 6
	*/
}

func main() {
	mux := NewHandler()
	http.ListenAndServe(":3000", mux)
}
