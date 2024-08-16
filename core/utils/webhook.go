package utils

import (
	"ctrl/core/config"
	"ctrl/core/updater"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strings"
	"time"
)

func SendWebHook(message []byte) bool {
	webhookUrl := config.Get().DiscordNotif.WebhookURL

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

func WebhookStatus() {
	message := fmt.Sprintf("CTRL Version: %s, system started on %s", updater.Version, time.Now().Format("2006-01-02 15:04:05"))
	res := SendWebHook([]byte(message))
	if !res {
		log.Warn().Msg("Failed to send webhook, verify webhook settings")
	}
}
