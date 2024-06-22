package main

import (
	qbit2 "ctrl/core/qbit"
	system "ctrl/core/system"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := "8080"
	result := fmt.Sprintf("0.0.0.0:%s", port)

	fmt.Println("Starting server on " + port)

	// ui

	// system power controls
	http.HandleFunc("/shutdown", system.ExecShutDown)
	http.HandleFunc("/reboot", system.ExecReboot)
	http.HandleFunc("/sleep", system.ExecSleep)
	// misc stuff
	http.HandleFunc("/status", status)
	http.HandleFunc("/test", test)
	//http.HandleFunc("/device", deviceCheck)

	// start periodic func
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
		go qbit2.CheckQBitStatus()
		go qbit2.SearchQbitStalled()
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
