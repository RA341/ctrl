package qbit

import (
	"ctrl/core/config"
	"ctrl/core/utils"
)

type Check int

const (
	ClientCheck Check = iota
	StalledCheck
)

func RunQbitChecks(checkFilter []Check) {
	auth := loginToQbit(config.Get().Qbit.User, config.Get().Qbit.Pass)

	if auth == "" {
		message := []byte("Error failed to login to Qbit")
		utils.SendWebHook(message)
		return
	}

	for _, checkItem := range checkFilter {
		runCheckFunc(checkItem, auth)
	}

}

func runCheckFunc(checkFilter Check, auth string) {
	switch checkFilter {
	case ClientCheck:
		checkClientStatus(auth)
	case StalledCheck:
		checkStalledTorrents(auth)
	}
}
