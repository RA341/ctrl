package qbit

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// loginToQbit login to qbit
func loginToQbit(username string, pass string) string {
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

// makeGetRequestToClient makes a get request with the specified path and returns json data
func makeGetRequestToClient(auth string, path string, isList bool) ([]map[string]interface{}, map[string]interface{}) {
	req, err := http.NewRequest("GET", path, strings.NewReader(""))
	if err != nil {
		log.Println("Failed to create request.Reason: " + err.Error())
		return emptyState(isList)
	}

	req.AddCookie(&http.Cookie{Name: "SID", Value: auth})

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Failed to send request\nReason: " + err.Error())
		return emptyState(isList)
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
			return emptyState(isList)
		}

		var data []map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Failed to unmarshal json", err)
			return emptyState(isList)
		}

		return data, nil

	} else {
		log.Println("Request failed: " + res.Status)
	}
	return emptyState(isList)
}

func emptyState(isList bool) ([]map[string]interface{}, map[string]interface{}) {
	if isList {
		return []map[string]interface{}{}, nil
	}
	return nil, map[string]interface{}{}

}
