package service

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/kamijin-fanta/sakuraio-cli/lib"
)

type DataStoreChannelOptions struct {
	Module    *string
	Size      *string
	Unit      *string
	Token     *string
	Order     *string
	Cursor    *string
	After     *string
	Before    *string
	Channel   *string
	Project   *string
	RawOutput *bool
}
type DataStoreMessagesOption struct {
	Module    *string
	Size      *string
	Order     *string
	Cursor    *string
	After     *string
	Before    *string
	Project   *string
	RawOutput *bool
	Token     *string
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
	paramSet(param, "size", *options.Size)
	paramSet(param, "unit", *options.Unit)
	paramSet(param, "order", *options.Order)
	paramSet(param, "cursor", *options.Cursor)
	paramSet(param, "after", *options.After)
	paramSet(param, "before", *options.Before)
	paramSet(param, "channel", *options.Channel)

	body, err := lib.HTTPGet("datastore/v1/channels?" + param.Encode())
	checkError("HTTP ERROR", err, body)
	if *options.RawOutput {
		fmt.Println(body)
		return
	}

	switch *options.Unit {
	case "channel":
		var channels ChannelsChannelResponse
		err = json.Unmarshal([]byte(body), &channels)
		checkError("JSON format error", err, body)

		meta := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
		fmt.Fprintln(meta, "Count\tMatch\tCursor")
		fmt.Fprintf(meta, "%d\t%d\t%s\n---\n", channels.Meta.Count, channels.Meta.match, channels.Meta.Cursor)
		meta.Flush()

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
		fmt.Fprintln(w, "Module\tCh\tType\tValue\tDatetime")
		for _, v := range channels.Results {
			fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\n", v.Module, v.Channel, v.Type, v.ValueStr, v.Datetime)
		}
		w.Flush()

	default: // message
		var messages ChannelsMessageResponse
		err = json.Unmarshal([]byte(body), &messages)
		checkError("JSON format error", err, body)

		meta := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
		fmt.Fprintln(meta, "Count\tMatch\tCursor")
		fmt.Fprintf(meta, "%d\t%d\t%s\n---\n", messages.Meta.Count, messages.Meta.match, messages.Meta.Cursor)
		meta.Flush()

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
		fmt.Fprintln(w, "Module\tDatetime\tChannelMessage(ch,type,value,datetime)")
		for _, v := range messages.Results {
			for i, c := range v.Channels {
				value := parseToString(c.Value)

				if i == 0 {
					fmt.Fprintf(w, "%s\t%s\t%d,%s,%s,%s\n", v.Module, v.Datetime, c.Channel, c.Type, value, c.Datetime)
				} else {
					fmt.Fprintf(w, "%s\t%s\t%d,%s,%s,%s\n", "", "", c.Channel, c.Type, value, c.Datetime)
				}

			}
		}
		w.Flush()
	}

}

func DataStoreMessagesCmd(options DataStoreMessagesOption) {
	token := GetToken(*options.Token, *options.Project, "datastore")

	param := url.Values{}
	param.Add("token", token)

	if *options.Module != "" {
		param.Add("module", *options.Module)
	}
	paramSet(param, "size", *options.Size)
	paramSet(param, "order", *options.Order)
	paramSet(param, "cursor", *options.Cursor)
	paramSet(param, "after", *options.After)
	paramSet(param, "before", *options.Before)

	body, err := lib.HTTPGet("datastore/v1/messages?" + param.Encode())
	checkError("HTTP ERROR", err, body)
	if *options.RawOutput {
		fmt.Println(body)
		return
	}

	var messages MessagesResponse

	err = json.Unmarshal([]byte(body), &messages)
	checkError("JSON format error", err, body)

	meta := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
	fmt.Fprintln(meta, "Count\tMatch\tCursor")
	fmt.Fprintf(meta, "%d\t%d\t%s\n---\n", messages.Meta.Count, messages.Meta.match, messages.Meta.Cursor)
	meta.Flush()

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
	fmt.Fprintln(w, "Module\tDatetime\tType\tPayload")
	for _, v := range messages.Results {
		payload, _ := json.Marshal(v.Payload)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", v.Module, v.Datetime, v.Type, payload)
	}
	w.Flush()

}

func parseToString(value interface{}) string {
	switch value.(type) {
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case string:
		return value.(string)
	default:
		return "***"
	}
}

func printProjects(projects []Project) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tProjectName")
	for _, v := range projects {
		fmt.Fprintf(w, "%d\t%s\n", v.ID, v.Name)
	}
	w.Flush()
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
	Channel  int
	Datetime string
	Module   string
	Type     string
	ValueNum float64 `json:"value_num"`
	ValueStr string  `json:"value_str"`
}

type ChannelsMessageResponse struct {
	Meta    Meta
	Results []ChannelsMessageResult
}

type ChannelsMessageResult struct {
	Module   string
	Datetime string
	Channels []ChannelsMessageResultPayload
}

type ChannelsMessageResultPayload struct {
	Channel  int
	Datetime string
	Type     string
	Value    interface{}
}

type MessagesResponse struct {
	Meta    Meta
	Results []MessagesResult
}
type MessagesResult struct {
	Module   string
	Datetime string
	Type     string
	Payload  interface{}
}

type Meta struct {
	Count  int
	Cursor string
	match  int
}
