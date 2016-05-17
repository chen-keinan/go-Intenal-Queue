package mqueue

import (
	"encoding/json"
	//	"fmt"
	"log"
	"net/http"
)

func IndexEventConsumer(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var readMsg bool
	var scannerEvent IndexRequest
	select {
	case scannerEvent = <-IndexEventQueue:
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
		msg, err := json.Marshal(scannerEvent)
		if err != nil {
			log.Printf(err.Error())
		}
		w.Write(msg)
	}
}
