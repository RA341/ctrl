package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func InitDocker() *client.Client {
	apiClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return apiClient
}

func DisposeDocker(client *client.Client) {
	err := client.Close()
	if err != nil {
		panic(err)
	}
}

func ListDocker(client *client.Client) []types.Container {
	containers, err := client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	return containers
}

func RestartContainer(client *client.Client, containerId string) bool {
	err := client.ContainerRestart(context.Background(), containerId, container.StopOptions{})
	if err != nil {
		fmt.Println(fmt.Sprintf("Error restarting container %s: %s\n", containerId, err))
		return false
	}
	return true
}

func StopContainer(client *client.Client, containerId string) bool {
	err := client.ContainerStop(context.Background(), containerId, container.StopOptions{})
	if err != nil {
		fmt.Println(fmt.Sprintf("Error restarting container %s: %s\n", containerId, err))
		return false
	}
	return true
}

func StartContainer(client *client.Client, containerId string) bool {
	err := client.ContainerStart(context.Background(), containerId, container.StartOptions{})
	if err != nil {
		fmt.Println(fmt.Sprintf("Error restarting container %s: %s\n", containerId, err))
		return false
	}
	return true
}

func GetContainerIdFromName(client *client.Client, containerName string) string {
	for _, ctr := range ListDocker(client) {
		for _, name := range ctr.Names {
			if name[1:] == containerName {
				return ctr.ID
			}
		}
	}
	return ""
}
