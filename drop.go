package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// HttpCreate creates a new drop via http api call
func (dr *Drop) HttpCreate() (resp *http.Response, err error) {
	var url string = DROP_POST

	if dr.FlowId != "" {
		url = fmt.Sprintf("%s/%s", DROP_POST, dr.FlowId)

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

	resp, err = flowHttpPostRequest(payload, url)

	return
}

// Create creates a new drop via http or websocket
func (dr *Drop) Create() (err error) {
	var resp *http.Response

	if !flowthings.Config.Websocket {
		resp, err = dr.HttpCreate()
		if err != nil {
			Logger.Error(err)
			return
		}
		defer resp.Body.Close()
	} else {
		// Do Websocket request TODO create anoter function
	}

	dropCreateResp := struct {
		Head ResponseHead
		Body Drop
	}{}

	json.NewDecoder(resp.Body).Decode(&dropCreateResp)
	if dropCreateResp.Head.Status != StatusResourceCreated {
		err = &dropCreateResp.Head
		return
	}

	*dr = dropCreateResp.Body

	return
}

// HttpDelete deletes a drop from flow via http
func (d *Drop) HttpDelete() (resp *http.Response, err error) {
	url := fmt.Sprintf("%s/%s", DROP_POST, d.FlowId)

	if d.Id != "" {
		url = fmt.Sprintf("%s/%s", url, d.Id)
	} else {
		Logger.Warning("Drop ID not set. Deleting all frops from flow.")
	}

	resp, err = flowHttpDeleteRequest(url)
	if err != nil {
		Logger.Error(err)
	}

	return
}

// Delete deletes a drop from a flow via http or websocket
// If Drop Id is not provided, all drops from flow will be deleted
func (d *Drop) Delete() (rh ResponseHead, err error) {
	deleteResp := struct {
		Head ResponseHead
	}{}
	var resp *http.Response

	if d.FlowId == "" {
		err = errors.New("FlowId not set")
		return
	}

	if !flowthings.Config.Websocket {
		resp, err = d.HttpDelete()
		if err != nil {
			Logger.Error(err)
			return
		}
		defer resp.Body.Close()
	}

	json.NewDecoder(resp.Body).Decode(&deleteResp)
	if deleteResp.Head.Status != StatusResourceDeleted {
		err = &deleteResp.Head
		return
	}
	rh = deleteResp.Head
	emptyDrop := Drop{}
	*d = emptyDrop

	return
}

// HttpRead reads a single drop from a specific flow via http api request
func (d *Drop) HttpRead() (resp *http.Response, err error) {
	url := fmt.Sprintf("%s/%s/%s", DROP_POST, d.FlowId, d.Id)
	resp, err = flowHttpGetRequest(url)
	return
}

// Read reads a single drop from a specific flow
func (d *Drop) Read() (err error) {
	if d.FlowId == "" || d.Id == "" {
		err = errors.New("Id and FlowId must be set")
		return
	}

	var resp *http.Response

	readResponse := struct {
		Head ResponseHead
		Body Drop
	}{}
	readResponse.Body = *d

	if !flowthings.Config.Websocket {
		resp, err = d.HttpRead()
		if err != nil {
			Logger.Error(err)
			return
		}
		defer resp.Body.Close()
	}

	json.NewDecoder(resp.Body).Decode(&readResponse)
	if readResponse.Head.Status != StatusRequestSuccessfull {
		err = &readResponse.Head
		return
	}

	*d = readResponse.Body

	return
}

/*
TODO
Change drop receiverst to Drop struct
remove return drops (use receiver struct)

Flow methods:
- TrackFrom
- TrackTo

Track methods:
- AddFlow
- RemoveFlow
*/
