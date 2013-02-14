package jsonutils_test

import (
	"github.com/bashtian/jsonutils"
	"log"
	"os"
)

func Example() {
	b := []byte(`{"intArray":[1,2],"floatArray":[1.12,2.23],"stringArray":["a","b"],"boolean":true,"null":null,"number":123,"float":123.12,"object":{"a":"b","c":"d","e":"f"},"string":"HelloWorld"}`)

	f, err := jsonutils.ParseJson(b)
	if err != nil {
		log.Fatal(err)
	}
	jsonutils.Writer = os.Stdout
	jsonutils.PrintGo(f)
	// Output:
	// type Data struct {
	// Boolean bool `json:"boolean"`
	// Float float64 `json:"float"`
	// FloatArray []float64 `json:"floatArray"`
	// IntArray []float64 `json:"intArray"`
	// Null interface{} `json:"null"`
	// Number int64 `json:"number"`
	// Object struct {
	// A string `json:"a"`
	// C string `json:"c"`
	// E string `json:"e"`
	// } `json:"object"`
	// String string `json:"string"`
	// StringArray []string `json:"stringArray"`
	// }

}
