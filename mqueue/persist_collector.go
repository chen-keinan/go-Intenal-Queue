package mqueue

import (
	//	"io/ioutil"
	"encoding/json"
	"fmt"
	"jfrog.com/xray/file"
	"log"
	"net/http"
)

// A buffered channel that we can send work requests on.
var PersistEventQueue = make(chan file.File, 100)

func PersistEventCollector(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// read message from request
	decoder := json.NewDecoder(r.Body)
	var file file.File
	err := decoder.Decode(&file)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	// read message from request
	select {
	case PersistEventQueue <- file:
		{
			log.Printf("persist request queued")
		}
	default:
		log.Printf("no message sent")
	}
	w.WriteHeader(http.StatusCreated)
}
