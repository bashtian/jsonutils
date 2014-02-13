package jsonutils_test

import (
	"log"

	"github.com/bashtian/jsonutils"
)

func Example() {
	b := []byte(`{"intArray":[1,2],"floatArray":[1.12,2.23],"stringArray":["a","b"],"boolean":true,"null":null,"number":123,"float":123.12,"object":{"a":"b","c":"d","e":"f"},"string":"HelloWorld"}`)

	f, err := jsonutils.ParseJson(b)
	if err != nil {
		log.Fatal(err)
	}
	jsonutils.PrintGo(f, "Example")
	// Output:
	// type Example struct {
	// Boolean bool `json:"boolean"` // true
	// Float float64 `json:"float"` // 123.12
	// FloatArray []float64 `json:"floatArray"` // 1.12
	// IntArray []float64 `json:"intArray"` // 1
	// Null interface{} `json:"null"` // <nil>
	// Number int64 `json:"number"` // 123
	// Object struct {
	// A string `json:"a"` // b
	// C string `json:"c"` // d
	// E string `json:"e"` // f
	// } `json:"object"`
	// String string `json:"string"` // HelloWorld
	// StringArray []string `json:"stringArray"` // a
	// }
}
