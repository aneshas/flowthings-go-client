package main

import (
	"fmt"
	"os"
)

type MyDrop struct {
	Foo    string
	Bar    string
	Nested map[string]string
}

var Logger ILogger

func init() {
	logger := DefaultLogger{}
	logger.Init()
	Logger = &logger
}

// TODO Flowthings cli
func main() {

	// Authenticate
	config := FlowConfig{
		Username:  os.Getenv("FT_USERNAME"),
		Token:     os.Getenv("FT_TOKEN"),
		Websocket: false,
	}

	Ft, err := NewFlowthings(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create drop
	dropRequest := new(DropRequest)

	location := Location{
		Lat: 87.89898989,
		Lon: 87.8989,
	}

	elems := MyDrop{
		Foo:    "Bar",
		Bar:    "Baz",
		Nested: make(map[string]string),
	}

	elems.Nested["nested1"] = "nested value"
	elems.Nested["nested2"] = "another nested value"

	dropRequest.Elems = elems
	dropRequest.FlowId = "f551d2c940cf213ccab26343d"
	dropRequest.Location = location
	dropRequest.Path = "/anes/otoka"

	drop, err := Ft.DropCreate(dropRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(drop)

	// Delete drop
	d := Drop{
		FlowId: "f551d2c940cf213ccab26343d",
		Id:     "d559d04375bb70963aca88045",
	}

	resp, err := Ft.DropDelete(&d)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("RESP")
	fmt.Println(resp)
}
