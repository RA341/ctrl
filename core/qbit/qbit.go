package qbit

import (
	"ctrl/core/config"
	"ctrl/core/utils"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/rs/zerolog/log"
)

type Check int

const (
	ClientCheck Check = iota
	StalledCheck
)

var qbitBasePath string

func InitBasePath() {
	qbitBasePath = fmt.Sprintf("%s/api/v2", config.Get().Qbit.Url)
}

func RunQbitChecks(checkFilter []Check, cli *client.Client) {
	url := qbitBasePath + loginPath

	auth := loginToQbit(url, config.Get().Qbit.User, config.Get().Qbit.Pass)

	if auth == "" {
		message := []byte("Error failed to login to Qbit")
		utils.SendWebHook(message)
		return
	}

	for _, checkItem := range checkFilter {
		runCheckFunc(checkItem, auth, cli)
	}

}

// Status logs into to qbit once to check if it is properly setup
func Status() {
	if !config.Get().Qbit.Enable {
		log.Debug().
			Bool("Qbit setting", config.Get().Qbit.Enable).
			Msg("Qbit is disabled")
		return
	}

	res := loginToQbit(config.Get().Qbit.Url, config.Get().Qbit.User, config.Get().Qbit.Pass)
	if res == "" {
		log.Warn().Msg("Qbit is not working, check the qbit config")
	}
}

func runCheckFunc(checkFilter Check, auth string, cli *client.Client) {
	switch checkFilter {
	case ClientCheck:
		checkClientStatus(auth, qbitBasePath+clientStatusPath, cli)
	case StalledCheck:
		checkStalledTorrents(auth)
	}
}
