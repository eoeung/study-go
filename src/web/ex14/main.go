package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/antage/eventsource"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

type Message struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

// 현재 멀티 쓰레드 환경이기 때문에 핸들러로 바로 등록하지 않고,
// 쓰레드의 Queue형태(한 줄)로 보내는 것이 낫다고 판단해서 생성
// → sendMessage가 되면 채널에 값을 집어넣어서 다른 쓰레드에서 이 채널에 접근해서 pop할 수 있도록 설정
var msgCh chan Message

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	sendMessage(name, msg)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("name")
	sendMessage("", fmt.Sprintf("add user: %s", username))
}

func leftUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	sendMessage("", fmt.Sprintf("left user: %s", username))
}

func sendMessage(name, msg string) {
	// send message to all clients
	// 채널에 데이터 넣는 방법 (채널 인스턴스를 향해 넣을 데이터를 선언)
	// msgCh에 값이 전달되면, 이 값을 처리하는 go processMsgCh(es) 실행
	msgCh <- Message{name, msg}
}

// 만든 채널인 msgCh에서 pop해서 EventSource로 보내는 역할
func processMsgCh(es eventsource.EventSource) {
	for msg := range msgCh {
		fmt.Println("### msg ::: ", msg)
		data, _ := json.Marshal(msg)
		fmt.Println("### string(data) ::: ", string(data))

		// 클라이언트와 연결된 EventSource(여기서는 "/stream"요청으로 연결된 es)로
		// 연결된 모든 클라이언트에게 메시지 전송
		es.SendEventMessage(string(data), "", strconv.Itoa(time.Now().Nanosecond()))
	}
}

func testtt() {
	for msg := range msgCh {
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@", msg)
	}
}

func main() {
	msgCh = make(chan Message) // Unbuffered Channel

	// EventSource
	es := eventsource.New(nil, nil) // EventSource 인스턴스 생성 // 이벤트소스 오픈
	defer es.Close()

	// goroutine으로 실행
	// main()에서 한 번만 실행됨
	// 런타임이 프로그램 시작 ~ 종료까지 goroutine을 관리 → 백그라운드에서 실행되고 있음
	// msgCh 채널에 값이 들어오면, 여기에 있는 goroutine(수신자)가 msgCh의 값을 받아서 처리해줌
	go processMsgCh(es)
	go testtt()

	// gorilla/pat 사용
	mux := pat.New()
	mux.Post("/messages", postMessageHandler)
	// /stream 경로로 들어온 클라이언트는 SSE를 통해 이벤트를 받을 수 있음
	// eventsource 인스턴스가 경로로 들어오는 요청을 처리
	// 연결을 유지하면서 지속적으로 클라이언트에게 이벤트 및 메시지 전송
	mux.Handle("/stream", es)
	mux.Post("/users", addUserHandler)
	mux.Delete("/users", leftUserHandler)

	// negroni 사용
	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
