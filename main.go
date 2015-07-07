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

// TODO Flowthings cli for testing
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
	dropRequest.FlowId = "939393099330"
	dropRequest.Location = location
	dropRequest.Path = "/anes/otoka"

	Ft.DropCreate(dropRequest)

	fmt.Println(Ft.SessionId)
}
