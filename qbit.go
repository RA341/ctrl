package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	qBitBasePath     = "http://localhost:8085/api/v2"
	loginPath        = "/auth/login"
	listTorrentsPath = "/torrents/info"
	webhookUrl       = "https://discord.com/api/webhooks/1223892724537753661/5VvQzM9chKTUYkAxit3ddAf__8s_dybIbBQ2sB33n7S7RHgn4OzQ27XgXZ0f2qbP0S7w"
)

const username = "Test Hook"
const avatar_url = "https://i.imgur.com/KEungv8.png"

func SearchQbitStalled() {
	auth := LoginToQbit("r334", "Thisismyqbitpasskey#1505")

	if auth == "" {
		message := []byte("Error failed to login to Qbit")
		SendWebHook(message)
		log.Println("[ERROR] failed to login to qbit")
		return
	}

	if auth == "" {
		log.Println("Auth is empty")
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

func SendWebHook(message []byte) bool {
	payload := strings.NewReader(string(message))

	req, err := http.NewRequest("POST", webhookUrl, payload)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Println("Failed to create request.Reason: " + err.Error())
		return false
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Failed to send request\nReason: " + err.Error())
		return false
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close buffer body")
			return
		}
	}(res.Body)

	if res.StatusCode < 300 {
		log.Print("Sent to webhook")
		return true
	} else {
		log.Println("Request failed: " + res.Status)
		return false
	}
}

func formatDuration(d time.Duration) string {
	days := int64(d.Hours() / 24)
	hours := int64(d.Hours()) % 24
	minutes := int64(d.Minutes()) % 60
	seconds := int64(d.Seconds()) % 60

	return fmt.Sprintf("%02d Days, %02d Hours, %02d Minutes, %02d Seconds", days, hours, minutes, seconds)
}

func formatDiscordMessage(stalled map[string]interface{}) (error, []byte) {
	discordContent := make(map[string]interface{})

	discordContent["avatar_url"] = avatar_url
	discordContent["username"] = username

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
			SendWebHook(message)
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
			SendWebHook(marshal)
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

// LoginToQbit login to qbit
func LoginToQbit(username string, pass string) string {
	url := qBitBasePath + loginPath

	payload := strings.NewReader(fmt.Sprintf("username=%s&password=%s", username, pass))

	req, err := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Println("Failed to create request.Reason: " + err.Error())
		return ""
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Failed to send request\nReason: " + err.Error())
		return ""
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close buffer body")
			return
		}
	}(res.Body)

	if res.StatusCode == 200 {

		cookies := res.Cookies()

		for _, cookie := range cookies {
			if cookie.Name == "SID" {
				return cookie.Value
			}
		}

		log.Println("Failed to get auth cookie. Reason: could not find the 'SID' cookie header")
	} else {
		log.Println("Request failed: " + res.Status)
	}
	return ""
}

// get all stalled torrents with last active greater than some time

func getStalledTorrents(auth string, filter string) []map[string]interface{} {

	url := qBitBasePath + listTorrentsPath + "?filter=" + filter

	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		log.Println("Failed to create request.Reason: " + err.Error())
		return []map[string]interface{}{}
	}

	req.AddCookie(&http.Cookie{Name: "SID", Value: auth})

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Failed to send request\nReason: " + err.Error())
		return []map[string]interface{}{}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close buffer body")
			return
		}
	}(res.Body)

	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Failed to read body", err)
			return []map[string]interface{}{}
		}

		var data []map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Failed to unmarshal json", err)
			return []map[string]interface{}{}
		}

		return data

	} else {
		log.Println("Request failed: " + res.Status)
	}
	return []map[string]interface{}{}
}
