package qbit

import (
	"ctrl/core/config"
	utils "ctrl/core/utils"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

var (
	qBitBasePath = fmt.Sprintf("%s/api/v2", config.Get().Qbit.Url)
)

const (
	loginPath        = "/auth/login"
	listTorrentsPath = "/torrents/info"
	clientStatusPath = "/transfer/info"
)

func SearchQbitStalled() {
	auth := LoginToQbit(config.Get().Qbit.User, config.Get().Qbit.Pass)

	if auth == "" {
		message := []byte("Error failed to login to Qbit")
		utils.SendWebHook(message)
		log.Println("[ERROR] failed to login to qbit")
		return
	}

	allFilters := []string{"stalled", "stalled_downloading"}
	thresholdTime := time.Hour * 2

	for _, filter := range allFilters {
		stalled := getStalledTorrents(auth, filter)
		// all stalled torrents with time added greater than 2 hours will be sent to webhook
		notifyStalledTorrents(stalled, thresholdTime)
	}

	// separate for downloading
	// because separate conditions that don't work with stalled
	stalled := getStalledTorrents(auth, "downloading")
	notifyDownloadingMetadataTorrents(stalled, thresholdTime)
}

func formatDiscordMessage(stalled map[string]interface{}) (error, []byte) {
	discordContent := make(map[string]interface{})

	discordContent["username"] = config.Get().DiscordNotif.Username
	discordContent["avatar_url"] = config.Get().DiscordNotif.AvatarURL

	message := fmt.Sprintf(
		"Stalled torrents detected.\nName: %s\nadded: %s ago\ncategory: %s\nStatus: %s", stalled["name"], formatDuration(timeSinceAdd(stalled)), stalled["category"], stalled["state"])
	discordContent["content"] = message

	marshal, err := json.Marshal(discordContent)
	if err != nil {
		log.Println("Failed to convert back to JSON")
		return nil, nil
	}
	return err, marshal
}

func formatDuration(d time.Duration) string {
	days := int64(d.Hours() / 24)
	hours := int64(d.Hours()) % 24
	minutes := int64(d.Minutes()) % 60
	seconds := int64(d.Seconds()) % 60

	return fmt.Sprintf("%02d Days, %02d Hours, %02d Minutes, %02d Seconds", days, hours, minutes, seconds)
}

func notifyStalledTorrents(stalled []map[string]interface{}, threshHold time.Duration) {
	for _, torrent := range stalled {
		duration := timeSinceAdd(torrent)

		// send web hook if torrent is stalled downloading
		// or if metadata has been downloading for more than an hour
		if duration.Hours() >= threshHold.Hours() && (torrent["state"] != "stalledUP") {
			if torrent["state"] != "metaDL" {
				continue
			}
			err, message := formatDiscordMessage(torrent)
			if err != nil {
				log.Println("Failed to format data")
			}
			utils.SendWebHook(message)
		} else {
			log.Println("[INFO] Torrent " + torrent["name"].(string) + "does not meet the criteria")
		}
	}
}

func notifyDownloadingMetadataTorrents(metaDown []map[string]interface{}, threshHold time.Duration) {
	for _, torrent := range metaDown {
		duration := timeSinceAdd(torrent)

		// send web hook if torrent is torrent downloading
		//but metadata down has been stuck for more then hour
		if duration.Hours() >= threshHold.Hours() && torrent["state"] == "metaDL" {
			err, marshal := formatDiscordMessage(torrent)
			if err != nil {
				log.Println("Failed to format data")
				continue
			}
			utils.SendWebHook(marshal)
		} else {
			log.Println("[INFO] Torrent " + torrent["name"].(string) + "does not meet the criteria")
		}
	}
}

func timeSinceAdd(torrent map[string]interface{}) time.Duration {
	timeSinceAdd := time.Unix(int64(torrent["added_on"].(float64)), 0)
	duration := time.Now().Sub(timeSinceAdd)
	return duration
}

// getStalledTorrents get all stalled torrents with last active greater than some time
func getStalledTorrents(auth string, filter string) []map[string]interface{} {
	url := qBitBasePath + listTorrentsPath + "?filter=" + filter
	res, _ := MakeGetRequestToClient(auth, url, true)
	return res
}
