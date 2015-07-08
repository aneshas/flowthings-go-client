// Flowthings.io Go package
// Manage drops, flows and tracks
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

const (
	WS_AUTH_URL              string = "https://ws.flowthings.io/session"
	WS_URL                   string = "wss://ws.flowthings.io/session/%s/ws"
	DROP_POST                string = "https://api.flowthings.io/v0.1/%s/drop"
	StatusRequestSuccessfull int    = 200
	StatusResourceUpdated    int    = 200
	StatusResourceCreated    int    = 201
	StatusBadRequest         int    = 400
	StatusUnauthorized       int    = 401
	StatusServiceUnavailable int    = 503
)

var DebugLevel int
var Logger ILogger

func openWebsocket(ft *Flowthings) {
	wsUrl := fmt.Sprintf(WS_URL, ft.SessionId)
	wsOrigin, _ := os.Hostname()

	ws, err := websocket.Dial(wsUrl, "", wsOrigin)
	if err != nil {
		Logger.Error(err)
		Logger.Info("Connection failed. Reconnecting...")
		return
	}
	Logger.Info("Websocket connection established.")
	ft.Ws = ws
}

func prepareHttpHeadersAndUrl(
	method string,
	url string,
	body io.Reader, ft *Flowthings) (req *http.Request, err error) {

	url = fmt.Sprintf(url, ft.Config.Username)

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		Logger.Error(err)
		return
	}

	req.Header.Add("X-Auth-Account", ft.Config.Username)
	req.Header.Add("X-Auth-Token", ft.Config.Token)
	req.Header.Add("Content-Type", "application/json")

	return
}

func flowHttpPostRequest(
	payload []byte,
	url string, ft *Flowthings) (resp *http.Response, err error) {

	httpClient := http.Client{}
	req, err := prepareHttpHeadersAndUrl("POST", url, bytes.NewBuffer(payload), ft)
	resp, err = httpClient.Do(req)

	return
}

func flowHttpDeleteRequest(url string, ft *Flowthings) (resp *http.Response, err error) {
	httpClient := http.Client{}
	req, err := prepareHttpHeadersAndUrl("DELETE", url, nil, ft)
	resp, err = httpClient.Do(req)

	return
}

// NewFlowthings creates new Flowthings struct used for further operations on flowthings primitives
func NewFlowthings(config FlowConfig) (ft *Flowthings, err error) {
	ft = new(Flowthings)
	ft.Config = new(FlowConfig)
	ft.Config = &config

	if !config.Websocket {
		return
	}

	req, err := http.NewRequest("POST", WS_AUTH_URL, nil)
	if err != nil {
		Logger.Error(err)
		return
	}

	req.Header.Add("X-Auth-Account", config.Username)
	req.Header.Add("X-Auth-Token", config.Token)

	httpClient := http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		Logger.Error(err)
		return
	}
	defer resp.Body.Close()

	response := new(AuthResponse)
	json.NewDecoder(resp.Body).Decode(&response)
	if response.Head.Status != StatusResourceCreated {
		err = &response.Head
		return
	}

	ft.SessionId = response.Body.Id
	openWebsocket(ft)

	return
}
