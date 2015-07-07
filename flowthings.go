// Flowthings.io Go package
// Manage your drops, flows and tracks
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

const (
	WS_AUTH_URL string = "https://ws.flowthings.io/session"
	WS_URL      string = "wss://ws.flowthings.io/session/%s/ws"

	DROP_POST string = "https://api.flowthings.io/v0.1/%s/drop/"
)

func openWebsocket(ft *Flowthings) {
	wsUrl := fmt.Sprintf(WS_URL, ft.SessionId)
	wsOrigin, _ := os.Hostname()

	fmt.Println(wsUrl)

	//for {
	ws, err := websocket.Dial(wsUrl, "", wsOrigin)
	if err != nil {
		log.Println(err)
		log.Println("Connection failed. Reconnecting...")
		return
		//continue
	}
	log.Println("Websocket connection established.")
	ft.Ws = ws
	//break
	//}
}

func flowHttpPostRequest(
	payload []byte,
	url string, ft *Flowthings) (resp *http.Response, err error) {

	httpClient := http.Client{}
	url = fmt.Sprintf(url, ft.Config.Username)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("X-Auth-Account", ft.Config.Username)
	req.Header.Add("X-Auth-Token", ft.Config.Token)
	req.Header.Add("Content-Type", "application/json")

	resp, err = httpClient.Do(req)

	return
}

func NewFlowthings(config FlowConfig) (ft *Flowthings, err error) {
	ft = new(Flowthings)
	ft.Config = new(FlowConfig)
	ft.Config = &config

	if !config.Websocket {

		return
	}

	req, err := http.NewRequest("POST", WS_AUTH_URL, nil)

	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("X-Auth-Account", config.Username)
	req.Header.Add("X-Auth-Token", config.Token)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	response := new(AuthResponse)
	json.NewDecoder(resp.Body).Decode(&response)

	if !response.Head.Ok {
		err = &response.Head
		return
	}

	ft.SessionId = response.Body.Id
	openWebsocket(ft)

	return
}
