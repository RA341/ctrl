package qbit

import (
	"ctrl/core/config"
	"ctrl/core/utils"
	"fmt"
	"log"
)

func CheckQBitStatus() {
	auth := LoginToQbit(config.Get().Qbit.User, config.Get().Qbit.Pass)

	if auth == "" {
		message := []byte("Error failed to login to Qbit")
		utils.SendWebHook(message)
		log.Println("[ERROR] failed to login to qbit")
		return
	}

	checkStatus(auth)
}

func checkStatus(auth string) {
	url := qBitBasePath + clientStatusPath
	_, data := MakeGetRequestToClient(auth, url, false)

	val, ok := data["connection_status"]

	if !ok {
		utils.SendWebHook([]byte("Failed to get Qbit client status"))
		return
	}

	if val != "connected" {
		utils.SendWebHook([]byte(fmt.Sprintf("Warning torrent client status: %s", val)))
	}
	return
}
