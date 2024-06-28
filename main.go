package main

import (
	"ctrl/core/config"
	qbit "ctrl/core/qbit"
	system "ctrl/core/system"
	"ctrl/core/updater"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Flavour: %s\n Version: %s\nCommit: %s\nBuild Date: %s\n", updater.Date,
		updater.Version, updater.Commit, updater.Date)

	if runtime.GOOS == "linux" && !verifyRootStatus() {
		// todo windows support
		log.Fatal("The program needs to run with root privileges")
	}

	config.Load()
	qbit.InitBasePath()

	system.RegisterLinuxService()

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

	settings := config.Get()
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

func verifyRootStatus() bool {
	return os.Geteuid() == 0
}
