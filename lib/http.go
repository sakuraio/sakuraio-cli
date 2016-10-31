package lib

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func HTTPGet(requestPath string) (string, error) {
	settings := GetSetting()
	url := settings.BaseURL + requestPath

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(settings.APIToken, settings.APISecret)

	return httpRequest(req)
}

func HTTPPost(requestPath string, payload string) (string, error) {
	settings := GetSetting()
	url := settings.BaseURL + requestPath

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(settings.APIToken, settings.APISecret)

	return httpRequest(req)
}

func HTTPDelete(requestPath string) (string, error) {
	settings := GetSetting()
	url := settings.BaseURL + requestPath

	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(settings.APIToken, settings.APISecret)

	return httpRequest(req)
}

func httpRequest(request *http.Request) (string, error) {
	client := &http.Client{}
	res, err := client.Do(request)

	if err != nil {
		return "", errors.New("[HTTP Client ERROR] " + err.Error())
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return string(body), errors.New("[ERROR HTTP Status] Status: " + res.Status)
	}

	return string(body), nil
}
