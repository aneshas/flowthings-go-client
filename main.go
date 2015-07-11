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

	_, err := NewFlowthings(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Flows */
	/*flow := Flow{
		Path:        "/anes/new_flow_from_go_plugin",
		Description: "Flow created with go plugin",
		Capacity:    1000,
	}

	err = flow.Create()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(flow)*/

	flow := Flow{}
	flow.Id = "f559fa0ce68056d07d5137534"
	//flow.Description = "Edited flow description"
	//flow.Capacity = 100

	// err = flow.Update()
	err = flow.Delete()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(flow)

	/* Drops */
	/*
		// Create drop

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

		drop := Drop{
			Elems:    elems,
			FlowId:   "f551d2c940cf213ccab26343d",
			Location: location,
			Path:     "/anes/otoka",
		}

		err = drop.Create()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(drop)

		// Update drop
		elems.Bar = "Updated bar, yaay :D"
		drop.Elems = elems
		drop.Update()

		// Delete drop
		d := Drop{
			FlowId: "f551d2c940cf213ccab26343d",
			Id:     "d559d04375bb70963aca88045",
		}

		resp, err := d.Delete()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(&resp)

		// Read drop
		dd := Drop{}
		dd.FlowId = "f551d2c940cf213ccab26343d"
		dd.Id = "d559e28a368056d2d0fc1c866"

		err = d.Read()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Drop read:")
		fmt.Println(d)
	*/

}
