package main

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/net/websocket"
)

type Flowthings struct {
	SessionId string
	Ws        *websocket.Conn
	Config    *FlowConfig
}

type FlowConfig struct {
	Username  string
	Token     string
	Websocket bool
}

type ResponseHead struct {
	Ok       bool
	Status   int
	Messages []string
	Errors   []string
}

func (rh *ResponseHead) Error() string {
	str := fmt.Sprintf("Status: %d", rh.Status)

	for _, err := range rh.Errors {
		str = fmt.Sprintf("%s\n%s\n", str, err)
	}

	for _, msg := range rh.Messages {
		str = fmt.Sprintf("%s\n%s\n", str, msg)
	}

	return str
}

type authResponseBody struct {
	Id string
}

type Location struct {
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Specifiers map[string]string `json:"specifiers,omitempty"`
}

type Drop struct {
	Id           string      `json:"id,omitempty"`
	FlowId       string      `json:"flowId,omitempty"`
	CreationDate int64       `json:"creationDate,omitempty"`
	Path         string      `json:"path,omitempty"`
	Location     Location    `json:"location,omitempty"`
	Fhash        string      `json:"fhash,omitempty"`
	Elems        interface{} `json:"elems,omitempty"`
}

func (d Drop) String() string {
	t := getTime(d.CreationDate)
	str := fmt.Sprintf("Id: %s\nFlowId: %s\nCreationDate: %s\nPath: %s",
		d.Id, d.FlowId, t, d.Path)

	return str
}

// TODO implement io.ReadWriter for flow and Stringer for all primitives
type Flow struct {
	Id           string `json:"id,omitempty"`
	Path         string `json:"path"`
	Description  string `json:"description,omitempty"`
	Filter       string `json:"filter,omitempty"`
	Capacity     int    `json:"capacity,omitempty"`
	CreationDate int64  `json:"creationDate,omitempty"`
	LastEditDate int64  `json:lastEditDate,omitempty`
}

func (f Flow) String() string {
	t := getTime(f.CreationDate)
	et := getTime(f.LastEditDate)

	return fmt.Sprintf(
		"Id: %s\nPath: %s\nDescription: %s\nFilter: %s\nCapacity: %d\nCreationDate: %s\nLasteEditDate: %s\n",
		f.Id, f.Path, f.Description, f.Filter, f.Capacity, t, et)
}

type Track struct {
	Id          string   `json:"id,omitempty"`
	Source      string   `json:"source"`
	Filter      string   `json:"filter,omitempty"`
	Js          string   `json:"js,omitempty"`
	Destination []string `json:"destination"`
}

func getTime(t int64) time.Time {
	timeStr := fmt.Sprintf("%d", t)
	msInt, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		Logger.Error(err)
	}
	return time.Unix(0, msInt*int64(time.Millisecond))
}
