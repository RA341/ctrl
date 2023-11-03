package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	port := "8080"
	result := fmt.Sprintf(":%s", port)

	http.HandleFunc("/shutdown", shutDownCmd)
	http.HandleFunc("/reboot", rebootCmd)
	http.HandleFunc("/status", status)
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

func shutDownCmd(_ http.ResponseWriter, _ *http.Request) {
	// Use the "shutdown" command to shut down the computer immediately
	cmd := exec.Command("sudo", "shutdown", "now")

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
