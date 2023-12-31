package main

import (
	"fmt"

	"github.com/countersoda/jsongo"
)

func main() {
	jsonString := `{"name": "John", "age": 30, "city": "New York"}`

	jsValue := jsongo.ParseFromString(jsonString)

	if jsValue != nil {
		fmt.Println(jsValue.ToString())
	} else {
		fmt.Println("Invalid json")
	}
}
