package main

import (
	"fmt"
	probing "github.com/prometheus-community/pro-bing"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	port := "8080"
	result := fmt.Sprintf(":%s", port)

	go periodicTask()

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

func deviceCheck(w http.ResponseWriter, _ *http.Request) {
	res := canPing("192.168.1.243")

	if res {
		_, err := w.Write([]byte("Device is up"))
		if err != nil {
			fmt.Println("Whoa there, something went wrong!" + err.Error())
			return
		}
	} else {
		_, err := w.Write([]byte("Device is down"))
		if err != nil {
			fmt.Println("Whoa there, something went wrong!" + err.Error())
			return
		}
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

func shutDown() *exec.Cmd {
	return exec.Command("sudo", "shutdown", "now")
}

func shutDownCmd(_ http.ResponseWriter, _ *http.Request) {
	// Use the "shutdown" command to shut down the computer immediately
	cmd := shutDown()
	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Whoa there, something went wrong!")
		println(err.Error())
	}
}

func rebootCmd(_ http.ResponseWriter, _ *http.Request) {
	// Use the "shutdown" command to shut down the computer immediately
	cmd := exec.Command("sudo", "reboot", "now")

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Whoa there, something went wrong!")
		println(err.Error())
	}
}

func periodicTask() {
	for {
		fmt.Println("Running Device check")

		res := canPing("192.168.1.243")
		if !res {
			fmt.Println("Device is down, shutting down")
			cmd := shutDown()
			err := cmd.Run()
			if err != nil {
				fmt.Println("Whoa there, something went wrong!")
				println(err.Error())
			}
		}
		fmt.Println("Device is up")
		time.Sleep(1 * time.Hour)
	}
}

func canPing(ipAddress string) bool {
	pinger, err := probing.NewPinger(ipAddress)
	pinger.SetPrivileged(true)

	if err != nil {
		fmt.Printf("Error creating pinger: %s\n", err)
		return false
	}

	pinger.Count = 4 // Number of ping packets to send

	err = pinger.Run()
	if err != nil {
		fmt.Printf("Error while pinging: %s\n", err)
		return false
	}

	stats := pinger.Statistics()
	return stats.PacketsRecv > 0
}
