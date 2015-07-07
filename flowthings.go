// Flowthings.io Go package
// Manage your drops, flows and tracks
package main

import (
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
)

func openWebsocket(ft *Flowthings) {
	wsUrl := fmt.Sprintf(WS_URL, ft.SessionId)
	wsOrigin, _ := os.Hostname()

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

func NewFlowthings(identity Identity) (ft *Flowthings, err error) {
	ft = new(Flowthings)

	req, err := http.NewRequest("POST", WS_AUTH_URL, nil)

	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("X-Auth-Account", identity.Username)
	req.Header.Add("X-Auth-Token", identity.Token)

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
