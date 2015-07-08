package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// HttpDropCreate creates a new drop via http api call
func (ft *Flowthings) HttpDropCreate(dr *DropRequest) (resp *http.Response, err error) {

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
		resp, err = ft.HttpDropCreate(dr)
		if err != nil {
			Logger.Error(err)
			return
		}
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

// HttpDropDelete deletes a drop from flow via http
func (ft *Flowthings) HttpDropDelete(d *Drop) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s/%s", DROP_POST, d.FlowId)

	if d.Id != "" {
		url = fmt.Sprintf("%s/%s", url, d.Id)
	} else {
		Logger.Warning("Drop ID not set. Deleting all frops from flow.")
	}

	resp, err = flowHttpDeleteRequest(url, ft)
	if err != nil {
		Logger.Error(err)
	}

	return
}

// DropDelete deletes a drop from a flow via http or websocket
// If Drop Id is not provided, all drops from flow will be deleted
func (ft *Flowthings) DropDelete(d *Drop) (rh ResponseHead, err error) {
	deleteResp := new(struct {
		Head ResponseHead
	})
	var resp *http.Response

	if d.FlowId == "" {
		err = errors.New("FlowId not set")
		return
	}

	if !ft.Config.Websocket {
		resp, err = ft.HttpDropDelete(d)
		if err != nil {
			Logger.Error(err)
			return
		}
		defer resp.Body.Close()
	}

	json.NewDecoder(resp.Body).Decode(&deleteResp)
	rh = deleteResp.Head

	return
}
