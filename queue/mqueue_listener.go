package main

import (
	"flag"
	"fmt"
	"jfrog.com/xray/mqueue"
	"net/http"
	"os"
)

func main() {
	// Parse the command-line flags.
	flag.Parse()
	qServer := os.Getenv("QSERVER")
	if len(qServer) == 0 {
		qServer = "127.0.0.1:8000"
	}
	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/scanner", mqueue.IndexEventCollector)
	http.HandleFunc("/scanner/local", mqueue.IndexEventConsumer)
	http.HandleFunc("/persist", mqueue.PersistEventCollector)
	http.HandleFunc("/persist/db", mqueue.PersistEventConsumer)
	http.HandleFunc("/ping", mqueue.Ping)
	// Start the HTTP server!
	fmt.Println("Message queue is listening on port:", qServer)
	if err := http.ListenAndServe(qServer, nil); err != nil {
		fmt.Println(err.Error())
	}
}
