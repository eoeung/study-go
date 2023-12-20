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
	msgCh <- Message{name, msg} // 채널에 데이터 넣는 방법 (채널 인스턴스를 향해 넣을 데이터를 선언)
}

func processMsgCh(es eventsource.EventSource) {
	for msg := range msgCh {
		fmt.Println("### msg ::: ", msg)
		data, _ := json.Marshal(msg)
		fmt.Println("### string(data) ::: ", string(data))
		es.SendEventMessage(string(data), "", strconv.Itoa(time.Now().Nanosecond()))
	}
}

func main() {
	msgCh = make(chan Message)

	// EventSource
	es := eventsource.New(nil, nil) // EventSource 인스턴스 생성 // 이벤트소스 오픈
	defer es.Close()

	go processMsgCh(es)

	// gorilla/pat 사용
	mux := pat.New()
	mux.Post("/messages", postMessageHandler)
	mux.Handle("/stream", es)
	mux.Post("/users", addUserHandler)
	mux.Delete("/users", leftUserHandler)

	// negroni 사용
	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
