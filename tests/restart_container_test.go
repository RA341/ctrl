package tests

import (
	docker "ctrl/core/docker"
	"fmt"
	"testing"
)

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestRestart(t *testing.T) {
	qbitName := "qbittorrent"

	cli := docker.InitDocker()
	defer docker.DisposeDocker(cli)

	id := docker.GetContainerIdFromName(cli, qbitName)

	if id != "f6a2acc07706081483a9524976760a29288c3928e136d6e6ec589a482d4fefcd" {
		erf := fmt.Sprintf("Id mismatch wanted %s, got %s", "f6a2acc07706081483a9524976760a29288c3928e136d6e6ec589a482d4fefcd", id)
		t.Fatalf(erf)
	}

	res := docker.RestartContainer(cli, qbitName)
	if !res {
		t.Fatalf("Failed to restart")
	}
}
