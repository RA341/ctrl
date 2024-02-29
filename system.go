package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

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
