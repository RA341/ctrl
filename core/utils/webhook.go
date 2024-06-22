package utils

import (
	"io"
	"log"
	"net/http"
	"strings"
)

const webhookUrl = "https://discord.com/api/webhooks/1223892724537753661/5VvQzM9chKTUYkAxit3ddAf__8s_dybIbBQ2sB33n7S7RHgn4OzQ27XgXZ0f2qbP0S7w"

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
