package main

import (
	"encoding/json"
	"fmt"
)

func (ft *Flowthings) DropCreate(
	flowId string,
	elems interface{}, location Location) (drop Drop, err error) {

	request, err := json.Marshal(elems)

	drop = Drop{}

	fmt.Println(string(request))

	return
}
