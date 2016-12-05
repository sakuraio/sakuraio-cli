package service

import (
	"fmt"
	"log"

	"time"

	"encoding/json"

	"github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type WebsocketListenOptions struct {
	Token     *string
	Project   *string
	RawOutput *bool
}

func WebsocketListenCommand(options WebsocketListenOptions) {
	token := GetToken(*options.Token, *options.Project, "websocket")

	url := "wss://api-dev.sakura.io/ws/v1/" + token
	origin := "https://api-dev.sakura.io/ws/v1/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		logrus.Fatal(err)
	}
	var msg = make([]byte, 4096)
	var n int
	for {
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}

		if true || *options.RawOutput {
			fmt.Println(string(msg[:n]))
		} else {
			var message MessagesResult
			err := json.Unmarshal(msg[:n], &message)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"error":  err,
					"recive": string(msg[:n]),
				}).Warn("Cannot decode message")
			} else {
				payloadJson, _ := json.Marshal(message.Payload)
				fmt.Printf("%s\t%s\t%s\t%s\n", message.Module, message.Type, message.Datetime, string(payloadJson))
			}
		}

		time.Sleep(30 * time.Millisecond)
	}
}
