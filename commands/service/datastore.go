package service

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"

	"github.com/sakuraio/sakuraio-cli/lib"
	"math"
	"strconv"
)

type DataStoreChannelOptions struct {
	Module      *string
	Size        *int
	Token       *string
	Order       *string
	Cursor      *string
	After       *string
	Before      *string
	Channel     *string
	Project     *string
	RawOutput   *bool
	NoRecursive *bool
	MaxReq      *int
	BatchSize   *int
}
type DataStoreMessagesOption struct {
	Module      *string
	Size        *int
	Order       *string
	Cursor      *string
	After       *string
	Before      *string
	Project     *string
	RawOutput   *bool
	Token       *string
	NoRecursive *bool
	MaxReq      *int
	BatchSize   *int
}

func paramSet(values url.Values, key string, value string) {
	if value != "" {
		values.Add(key, value)
	}
}

func DataStoreChannelsCommand(options DataStoreChannelOptions) {
	token := GetToken(*options.Token, *options.Project, "datastore")

	param := url.Values{}
	param.Add("token", token)

	if *options.Module != "" {
		param.Add("module", *options.Module)
	}
	paramSet(param, "size", strconv.Itoa(*options.BatchSize))
	paramSet(param, "order", *options.Order)
	paramSet(param, "cursor", *options.Cursor)
	paramSet(param, "after", *options.After)
	paramSet(param, "before", *options.Before)
	paramSet(param, "channel", *options.Channel)

	tabWriter := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
	lastCursor := ""

	for i := 0; i < *options.MaxReq; i++ {
		if len(lastCursor) != 0 {
			param.Set("cursor", lastCursor)
		}
		body, err := lib.HTTPGet("datastore/v1/channels?" + param.Encode())
		checkError("HTTP ERROR", err, body)

		var channels ChannelsChannelResponse
		err = json.Unmarshal([]byte(body), &channels)
		checkError("JSON format error", err, body)

		lastCursor = channels.Meta.Cursor
		reamingCount := *options.Size

		if *options.NoRecursive && *options.RawOutput {
			fmt.Println(body)
		} else if *options.RawOutput {
			for _, row := range channels.Results {
				if reamingCount > 0 {
					str, _ := json.Marshal(row)
					fmt.Println(string(str))
				}
				reamingCount--
			}
		} else { // table output
			if !*options.NoRecursive {
				meta := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
				fmt.Fprintln(meta, "Count\tMatch\tCursor")
				fmt.Fprintf(meta, "%d\t%d\t%s\n---\n", channels.Meta.Count, channels.Meta.Match, channels.Meta.Cursor)
				meta.Flush()
			}

			if i == 0 {
				fmt.Fprintln(tabWriter, "Module\tCh\tType\tValue\tDatetime")
			}
			for _, v := range channels.Results {
				if reamingCount > 0 {
					fmt.Fprintf(tabWriter, "%s\t%d\t%s\t%s\t%s\n", v.Module, v.Channel, v.Type, v.ValueStr, v.Datetime)
				}
				reamingCount--
			}
			tabWriter.Flush()
		}

		isLastLoop := channels.Meta.Match == 0
		if *options.NoRecursive || isLastLoop || reamingCount <= 0 {
			break
		}
	}
}

func DataStoreMessagesCmd(options DataStoreMessagesOption) {
	token := GetToken(*options.Token, *options.Project, "datastore")

	param := url.Values{}
	param.Add("token", token)

	if *options.Module != "" {
		param.Add("module", *options.Module)
	}
	paramSet(param, "size", strconv.Itoa(*options.BatchSize))
	paramSet(param, "order", *options.Order)
	paramSet(param, "cursor", *options.Cursor)
	paramSet(param, "after", *options.After)
	paramSet(param, "before", *options.Before)

	// use when recursive fetch with cursor
	reamingCount := *options.Size
	if reamingCount == 0 {
		reamingCount = math.MaxInt32
	}
	lastCursor := ""
	tabWriter := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
	for i := 0; i < *options.MaxReq; i++ {
		if len(lastCursor) != 0 {
			param.Set("cursor", lastCursor)
		}

		body, err := lib.HTTPGet("datastore/v1/messages?" + param.Encode())
		checkError("HTTP ERROR", err, body)

		var messages MessagesResponse
		err = json.Unmarshal([]byte(body), &messages)
		checkError("JSON format error", err, body)

		lastCursor = messages.Meta.Cursor
		if !*options.NoRecursive && *options.RawOutput {
			// raw json recursive output in recursive
			for _, row := range messages.Results {
				if reamingCount > 0 {
					str, _ := json.Marshal(row)
					fmt.Println(string(str))
				}
				reamingCount--
			}
		} else if *options.RawOutput {
			// raw json output
			fmt.Println(body)
		} else {
			// table layout
			if *options.NoRecursive {
				meta := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
				fmt.Fprintln(meta, "Count\tMatch\tCursor")
				fmt.Fprintf(meta, "%d\t%d\t%s\n---\n", messages.Meta.Count, messages.Meta.Match, messages.Meta.Cursor)
				meta.Flush()
			}

			if i == 0 {
				fmt.Fprintln(tabWriter, "Module\tDatetime\tType\tPayload")
			}
			for _, v := range messages.Results {
				if reamingCount > 0 {
					payload, _ := json.Marshal(v.Payload)
					fmt.Fprintf(tabWriter, "%s\t%s\t%s\t%s\n", v.Module, v.Datetime, v.Type, payload)
				}
				reamingCount--
			}
			tabWriter.Flush()
		}
		isLastLoop := messages.Meta.Match == 0
		if *options.NoRecursive || isLastLoop || reamingCount <= 0 {
			break
		}
	}
}

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ChannelsChannelResponse struct {
	Meta    Meta
	Results []ChannelsChannelResult
}

type ChannelsChannelResult struct {
	Channel  int     `json:"channel"`
	Datetime string  `json:"datetime"`
	Module   string  `json:"module"`
	Type     string  `json:"type"`
	ValueNum float64 `json:"value_num"`
	ValueStr string  `json:"value_str"`
}

type ChannelsMessageResponse struct {
	Meta    Meta
	Results []ChannelsMessageResult
}

type ChannelsMessageResult struct {
	Id       string                         `json:"id"`
	Module   string                         `json:"module"`
	Datetime string                         `json:"datetime"`
	Channels []ChannelsMessageResultPayload `json:"channels"`
}

type ChannelsMessageResultPayload struct {
	Channel  int         `json:"channel"`
	Datetime string      `json:"datetime"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
}

type MessagesResponse struct {
	Meta    Meta
	Results []MessagesResult
}
type MessagesResult struct {
	Id       string      `json:"id"`
	Module   string      `json:"module"`
	Datetime string      `json:"datetime"`
	Type     string      `json:"type"`
	Payload  interface{} `json:"payload"`
}

type Meta struct {
	Count  int
	Cursor string
	Match  int
}
