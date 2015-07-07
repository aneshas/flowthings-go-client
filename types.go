package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Flowthings struct {
	SessionId string
	Ws        *websocket.Conn
	//Http
}

type Identity struct {
	Username string
	Token    string
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
	Lat        string
	Lon        string
	Specifiers map[string]string
}

type Drop struct {
	Id           string
	FlowId       string `json:"flowId"`
	CreationDate string `json:"creationDate"`
	Path         string
}

type Flow struct{}
type Track struct{}
