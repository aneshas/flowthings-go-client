package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Create creates a new flow via http api call
func (f *Flow) HttpCreate() (resp *http.Response, err error) {
	var url string = FLOW_POST
	var method string = "POST"

	if f.Id == "" && f.Path == "" {
		err = errors.New("Flow path is required")
		return
	}

	payload, err := json.Marshal(f)
	if err != nil {
		Logger.Error(err)
		return
	}

	if f.Id != "" {
		method = "PUT"
		url = fmt.Sprintf("%s/%s", url, f.Id)
	}
	resp, err = flowHttpRequest(method, payload, url)
	return
}

// Create creates a new flow
func (f *Flow) Create() (err error) {
	var resp *http.Response

	if !flowthings.Config.Websocket {
		resp, err = f.HttpCreate()
		if err != nil {
			Logger.Error(err)
			return
		}
		defer resp.Body.Close()
	}

	flowCreateResp := struct {
		Head ResponseHead
		Body Flow
	}{}

	json.NewDecoder(resp.Body).Decode(&flowCreateResp)
	if flowCreateResp.Head.Status != StatusResourceCreated &&
		flowCreateResp.Head.Status != StatusResourceUpdated {
		err = &flowCreateResp.Head
		return
	}

	*f = flowCreateResp.Body
	return
}

// Read reads a specific flow
func (f *Flow) Read() (err error) {
	return
}

// Update updates specific flow
func (f *Flow) Update() (err error) {
	return f.Create()
}

// HttpDelete removes a flow via http call
func (f *Flow) HttpDelete() (resp *http.Response, err error) {
	url := fmt.Sprintf("%s/%s", FLOW_POST, f.Id)

	if f.Id == "" {
		err = errors.New("Flow Id not set")
		return
	}

	resp, err = flowHttpDeleteRequest(url)
	if err != nil {
		Logger.Error(err)
	}
	return
}

// Delete removes a flow
func (f *Flow) Delete() (err error) {
	var resp *http.Response

	if !flowthings.Config.Websocket {
		resp, err = f.HttpDelete()
		if err != nil {
			Logger.Error(err)
			return
		}
		defer resp.Body.Close()
	}

	deleteResp := struct {
		Head ResponseHead
	}{}
	json.NewDecoder(resp.Body).Decode(deleteResp)
	if deleteResp.Head.Status != StatusResourceDeleted {
		err = &deleteResp.Head
		return
	}
	emptyFlow := Flow{}
	*f = emptyFlow

	return
}

// TrackTo creates a new track to the specified flows
func (f *Flow) TrackTo(flows ...Flow) (t Track, err error) {
	return
}

// Children returns children flows
func (f *Flow) Children() (flows []Flow) {
	return
}
