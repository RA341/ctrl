package qbit

import (
	"ctrl/core/config"
	"ctrl/core/docker"
	"ctrl/core/utils"
	"fmt"
	"github.com/docker/docker/client"
)

func checkClientStatus(auth string, url string, cli *client.Client) {
	_, data := makeGetRequestToClient(auth, url, false)

	val, ok := data["connection_status"]

	if !ok {
		utils.SendWebHook([]byte("Failed to get Qbit client status, restarting container"))
		return
	}

	if val != "connected" {
		utils.SendWebHook([]byte(fmt.Sprintf("Warning torrent client status: %s", val)))
		docker.RestartContainer(cli, docker.GetContainerIdFromName(cli, config.Get().Qbit.ContainerName))
	}
	return
}
