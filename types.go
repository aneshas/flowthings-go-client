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

type AuthResponse struct {
	Head responseHead
	Body authResponseBody
}

type responseHead struct {
	Ok       bool
	Status   int
	Messages []string
	Errors   []string
}

type DropRequest struct {
	Path     string      `json:"path,omitempty"`
	FlowId   string      `json:"flowId,omitempty"`
	Location Location    `json:"location,omitempty"`
	Elems    interface{} `json:"elems"`
}

func (rh *responseHead) Error() string {
	var str string
	str = fmt.Sprintf("Error code: %d", rh.Status)

	for _, err := range rh.Errors {
		str = fmt.Sprintf("%s\n%s\n", str, err)
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

type DropCreateResponse struct {
	Head responseHead
	Body Drop
}

type Drop struct {
	Id           string   `json:"id"`
	FlowId       string   `json:"flowId"`
	CreationDate int64    `json:"creationDate"`
	Path         string   `json:"path"`
	Location     Location `json:"location"`
}

func (d Drop) String() string {
	timeStr := fmt.Sprintf("%d", d.CreationDate)
	msInt, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		Logger.Error(err)
	}
	t := time.Unix(0, msInt*int64(time.Millisecond))

	str := fmt.Sprintf("Id: %s\nFlowId: %s\nCreationDate: %s\nPath: %s",
		d.Id, d.FlowId, t, d.Path)

	return str
}

// TODO implement io.ReadWriter for flow and Stringer for all primitives
type Flow struct{}
type Track struct{}
