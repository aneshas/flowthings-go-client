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
	identity := Identity{
		Username: os.Getenv("FT_USERNAME"),
		Token:    os.Getenv("FT_TOKEN"),
	}

	Ft, err := NewFlowthings(identity)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Create drop
	location := Location{
		Lat: "87.89898989",
		Lon: "87.8989",
	}

	elems := MyDrop{
		Foo:    "Bar",
		Bar:    "Baz",
		Nested: make(map[string]string),
	}

	elems.Nested["nested1"] = "nested value"
	elems.Nested["nested2"] = "another nested value"

	Ft.DropCreate("21092933", elems, location)

	fmt.Println(Ft.SessionId)
}
