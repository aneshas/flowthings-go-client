package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HttpCreate creates a new drop via http api call
func (ft *Flowthings) HttpCreate(dr *DropRequest) (resp *http.Response, err error) {

	var url string = DROP_POST

	if dr.FlowId != "" {
		url = fmt.Sprintf("%s/%s", DROP_POST, dr.FlowId)

		Logger.Info(url)

		if dr.Path != "" {
			dr.Path = ""
			Logger.Warning("Both Drop.Path and Drop.FlowId are set, Drop.FlowId will be used")
		}
	}

	payload, err := json.Marshal(dr)
	if err != nil {
		Logger.Error(err)
		return
	}

	resp, err = flowHttpPostRequest(payload, DROP_POST, ft)

	return
}

// DropCreate creates a new drop via http or websocket
func (ft *Flowthings) DropCreate(dr *DropRequest) (drop Drop, err error) {
	var resp *http.Response

	if !ft.Config.Websocket {
		// Do Http request
		resp, err = ft.HttpCreate(dr)
		defer resp.Body.Close()
	} else {
		// Do Websocket request TODO create anoter function
	}

	dropCreateResp := DropCreateResponse{}
	json.NewDecoder(resp.Body).Decode(&dropCreateResp)

	if dropCreateResp.Head.Status != StatusResourceCreated {
		err = &dropCreateResp.Head
		return
	}

	drop = dropCreateResp.Body

	return
}

// DropDelete deletes a drop from a flow
func (ft *Flowthings) DropDelete(d Drop) error {

	return nil
}
