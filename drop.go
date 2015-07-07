package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type dropCreateReq struct {
}

func (ft *Flowthings) HttpCreate(dr *DropRequest) (resp *http.Response, err error) {

	// TODO create json request and call flowHttpPostRequest
	var flowId string

	if dr.FlowId != "" {
		flowId = dr.FlowId
		dr.FlowId = ""
	}

	payload, err := json.Marshal(dr)
	if err != nil {
		log.Println(err)
		return
	}

	// TODO add it to url if present and path not empty
	fmt.Println(string(flowId))

	resp, err = flowHttpPostRequest(payload, DROP_POST, ft)

	return
}

func (ft *Flowthings) DropCreate(dr *DropRequest) (drop Drop, err error) {

	var resp *http.Response

	// Do Http request
	if !ft.Config.Websocket {
		resp, err = ft.HttpCreate(dr)
		defer resp.Body.Close()
	} else {
		// Do Websocket request TODO create anoter function
	}

	// TODO decode response
	drop = Drop{}

	fmt.Println("Drop created")
	fmt.Println(resp)

	return
}
