package mqueue

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// A buffered channel that we can send work requests on.
var IndexEventQueue = make(chan IndexRequest, 100)

func IndexEventCollector(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// read message from request
	decoder := json.NewDecoder(r.Body)
	var workEvent IndexRequest
	err := decoder.Decode(&workEvent)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	// add message to queue
	select {
	case IndexEventQueue <- workEvent:
		{
			msg, _ := ioutil.ReadAll(r.Body)
			logrus.Info("scanner request queued" + string(msg))
		}
	default:
		//	logrus.Info("no message sent")
	}
	w.WriteHeader(http.StatusCreated)
}
