package service

import (
	"encoding/json"

	"../"
)

func OutgoingWebhookAddCommand(project int, name string, url string, secret string) {
	regist := OutgoingWebhookRegistInfo{
		Name:    name,
		Type:    "outgoing-webhook",
		Project: project,
		URL:     url,
		Secret:  secret,
	}
	json, _ := json.Marshal(regist)

	commands.AddService(json)
}

type OutgoingWebhookRegistInfo struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Project int    `json:"project"`
	URL     string `json:"url"`
	Secret  string `json:"secret"`
}
