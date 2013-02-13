package jsonutils_test

import (
	"github.com/bashtian/jsonutils"
	"log"
	"os"
)

func Example() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	f, err := jsonutils.ParseJson(b)
	if err != nil {
		log.Fatal(err)
	}
	jsonutils.PrintGo(f)
}
