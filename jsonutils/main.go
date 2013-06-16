package main

import (
	"flag"
	"github.com/bashtian/jsonutils"
	"io/ioutil"
	"log"
	"strings"
)

var filename = flag.String("f", "", "use a file as input")

func main() {
	flag.Parse()
	url := flag.Arg(0)

	if *filename != "" {
		b, err := ioutil.ReadFile(*filename)
		if err != nil {
			log.Fatal(err)
		}
		f, err := jsonutils.ParseJson(b)
		if err != nil {
			log.Fatal(err)
		}
		s := strings.Split(*filename, ".")

		m := jsonutils.New(f, s[0])
		m.WriteGo()
	} else if url != "" {
		m, err := jsonutils.GetModel(flag.Arg(0))
		if err != nil {
			log.Fatal(err.Error())
		}
		m.WriteGo()
	} else {
		log.Fatal("missing URL")
	}

}
