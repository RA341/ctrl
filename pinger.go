package main

import (
	"fmt"
	probing "github.com/prometheus-community/pro-bing"
	"net/http"
)

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
