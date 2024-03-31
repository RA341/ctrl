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
	result := fmt.Sprintf("0.0.0.0:%s", port)

	fmt.Println("Starting server on " + port)

	http.HandleFunc("/shutdown", shutDownCmd)
	http.HandleFunc("/reboot", rebootCmd)
	http.HandleFunc("/status", status)
	http.HandleFunc("/device", deviceCheck)
	http.HandleFunc("/test", test)

	go runPeriodicTasks()

	err := http.ListenAndServe(result, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func runPeriodicTasks() {
	ticker := time.NewTicker(time.Hour * 1)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Running stalled torrent search")
		go SearchQbitStalled()
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
