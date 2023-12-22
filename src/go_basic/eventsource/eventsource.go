package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/antage/eventsource"
)

func main() {
	es := eventsource.New(nil, nil)
	defer es.Close()
	http.Handle("/events", es)
	go func() {
		id := 1
		for {
			es.SendEventMessage("tick", "tick-event", strconv.Itoa(id))
			id++
			time.Sleep(2 * time.Second)
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
