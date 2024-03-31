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

	fmt.Println("Starting server on " + port)

	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Running stalled torrent search")
		go SearchQbitStalled()
	}

	http.HandleFunc("/shutdown", shutDownCmd)
	http.HandleFunc("/reboot", rebootCmd)
	http.HandleFunc("/status", status)
	http.HandleFunc("/device", deviceCheck)
	http.HandleFunc("/test", test)

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
