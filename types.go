package main

import (
	"fmt"
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

type Drop struct {
	Id           string
	FlowId       string `json:"flowId"`
	CreationDate string `json:"creationDate"`
	Path         string
}

type Flow struct{}
type Track struct{}
