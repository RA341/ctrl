package main

import (
	"ctrl/core/config"
	qbit "ctrl/core/qbit"
	system "ctrl/core/system"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	config.Load()
	qbit.InitBasePath()

	settings := config.Get()

	// TODO ui

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

	port := strconv.Itoa(settings.Network.Port)
	result := fmt.Sprintf("%s:%s", settings.Network.Host, port)
	fmt.Println("Starting server on " + port)

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
		go qbit.RunQbitChecks([]qbit.Check{qbit.ClientCheck, qbit.StalledCheck})
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
