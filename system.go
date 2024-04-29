package main

import (
	"net/http"
)

func ExecShutDown(_ http.ResponseWriter, _ *http.Request) {
	// Use the "shutdown" command to shut down the computer immediately
	os := getOS()
	var cmds []string
	if os == "windows" {
		cmds = []string{"psshutdown", "-s", " -t", "0"}
	} else {
		cmds = []string{"sudo", "shutdown", "now"}
	}
	execShell(cmds)
}

func ExecReboot(_ http.ResponseWriter, _ *http.Request) {
	os := getOS()
	var cmds []string
	if os == "windows" {
		cmds = []string{"psshutdown", "-r", " -t", "0"}
	} else {
		cmds = []string{"sudo", "reboot", "now"}
	}
	execShell(cmds)
}

func ExecSleep(_ http.ResponseWriter, _ *http.Request) {
	os := getOS()
	var cmds []string
	if os == "windows" {
		cmds = []string{"psshutdown", "-d", " -t", "0"}
	} else {
		cmds = []string{"sudo", "sleep", "now"}
	}
	execShell(cmds)

}
