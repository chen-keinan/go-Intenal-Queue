package mqueue

import (
	"encoding/json"
	//	"fmt"
	"jfrog.com/xray/file"
	"log"
	"net/http"
)

func PersistEventConsumer(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var readMsg bool
	var file file.File
	select {
	case file = <-PersistEventQueue:
		{
			readMsg = true
		}
	default:
		{
			//	fmt.Println("no event found")
			readMsg = false
		}
	}
	// consume the event from queue.
	if readMsg {
		msg, err := json.Marshal(file)
		if err != nil {
			log.Printf(err.Error())
		}
		w.Write(msg)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("pong"))
}
