package utils

import (
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strings"
)

const webhookUrl = "https://discord.com/api/webhooks/1223892724537753661/5VvQzM9chKTUYkAxit3ddAf__8s_dybIbBQ2sB33n7S7RHgn4OzQ27XgXZ0f2qbP0S7w"

func SendWebHook(message []byte) bool {
	payload := strings.NewReader(string(message))

	req, err := http.NewRequest("POST", webhookUrl, payload)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Error().Err(err).Msg("Failed to create request")
		return false
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Error().Err(err).Msg("Failed to send request")
		return false
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close body")
			return
		}
	}(res.Body)

	if res.StatusCode < 300 {
		log.Debug().Msgf("Webhook sent: %s", res.Status)
		return true
	} else {
		log.Error().Msgf("Request failed: %s", res.Status)
		return false
	}
}
