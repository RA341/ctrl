package main

import (
	"fmt"
	"net/http"
)

var running = false

func main() {

	port := "8080"
	result := fmt.Sprintf(":%s", port)

	// Add the number of tasks to the WaitGroup counter
	wg.Add(1)

	// run sync commands
	periodicRcloneTask()

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
