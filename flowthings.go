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
	WS_AUTH_URL string = "https://ws.flowthings.io/session"
	WS_URL      string = "wss://ws.flowthings.io/session/%s/ws"
	DROP_POST   string = "https://api.flowthings.io/v0.1/%s/drop"
	FLOW_POST   string = "https://api.flowthings.io/v0.1/%s/flow"

	StatusResourceDeleted    int = 200
	StatusRequestSuccessfull int = 200
	StatusResourceUpdated    int = 200
	StatusResourceCreated    int = 201
	StatusBadRequest         int = 400
	StatusUnauthorized       int = 401
	StatusServiceUnavailable int = 503
)

var Debug bool
var Logger ILogger
var flowthings *Flowthings

func openWebsocket() {
	wsUrl := fmt.Sprintf(WS_URL, flowthings.SessionId)
	wsOrigin, _ := os.Hostname()

	ws, err := websocket.Dial(wsUrl, "", wsOrigin)
	if err != nil {
		Logger.Error(err)
		Logger.Info("Connection failed. Reconnecting...")
		return
	}
	Logger.Info("Websocket connection established.")
	flowthings.Ws = ws
}

func prepareHttpHeadersAndUrl(
	method string,
	url string,
	body io.Reader) (req *http.Request, err error) {

	url = fmt.Sprintf(url, flowthings.Config.Username)

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		Logger.Error(err)
		return
	}

	req.Header.Add("X-Auth-Account", flowthings.Config.Username)
	req.Header.Add("X-Auth-Token", flowthings.Config.Token)
	req.Header.Add("Content-Type", "application/json")

	return
}

func flowHttpGetRequest(url string) (resp *http.Response, err error) {
	httpClient := http.Client{}
	req, err := prepareHttpHeadersAndUrl("GET", url, nil)
	resp, err = httpClient.Do(req)

	return
}

func flowHttpRequest(
	method string,
	payload []byte,
	url string) (resp *http.Response, err error) {

	httpClient := http.Client{}

	req, err := prepareHttpHeadersAndUrl(
		method,
		url,
		bytes.NewBuffer(payload))

	resp, err = httpClient.Do(req)

	return
}

func flowHttpDeleteRequest(url string) (resp *http.Response, err error) {
	httpClient := http.Client{}
	req, err := prepareHttpHeadersAndUrl("DELETE", url, nil)
	resp, err = httpClient.Do(req)

	return
}

// NewFlowthings creates new Flowthings struct used for further operations on flowthings primitives
func NewFlowthings(config FlowConfig) (ft *Flowthings, err error) {
	ft = new(Flowthings)
	ft.Config = new(FlowConfig)
	ft.Config = &config
	flowthings = ft

	if !config.Websocket {
		return
	}

	// TODO Refactor: use flowHttpRequest
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

	response := struct {
		Head ResponseHead
		Body authResponseBody
	}{}
	json.NewDecoder(resp.Body).Decode(&response)
	if response.Head.Status != StatusResourceCreated {
		err = &response.Head
		return
	}

	ft.SessionId = response.Body.Id
	openWebsocket()

	return
}
