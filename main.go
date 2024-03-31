package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var running = false

func main() {

	port := "8080"
	result := fmt.Sprintf(":%s", port)

	http.HandleFunc("/shutdown", shutDownCmd)
	http.HandleFunc("/reboot", rebootCmd)
	http.HandleFunc("/status", status)
	http.HandleFunc("/device", deviceCheck)
	http.HandleFunc("/test", test)

	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Running stalled torrent search")
		go SearchQbitStalled()
	}

	err := http.ListenAndServe(result, nil)

	if err != nil {
		return
	}
}

func test(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		return
	}
}

func status(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
