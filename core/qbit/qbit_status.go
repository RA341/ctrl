package qbit

import (
	"ctrl/core/utils"
	"fmt"
)

func checkClientStatus(auth string) {
	url := qBitBasePath + clientStatusPath
	_, data := makeGetRequestToClient(auth, url, false)

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
