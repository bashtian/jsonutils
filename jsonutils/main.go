package main

import (
	"flag"
	"github.com/bashtian/jsonutils"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var filename = flag.String("f", "", "use a file as input")
var withExamples = flag.Bool("x", false, "print examples as comment")
var asJava = flag.Bool("j", false, "print Java instead of Go code")

func main() {
	flag.Parse()
	url := flag.Arg(0)

	var m *jsonutils.Model

	if *filename != "" {
		b, err := ioutil.ReadFile(*filename)
		if err != nil {
			log.Fatal(err)
		}
		s := strings.Split(*filename, ".")
		m, err = jsonutils.FromBytes(b, s[0])
		if err != nil {
			log.Fatal(err)
		}
	} else if url != "" {
		var err error
		m, err = jsonutils.GetModel(url)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		m, err = jsonutils.FromBytes(b, "")
		if err != nil {
			log.Fatal(err)
		}
	}

	m.WithExample = *withExamples
	if *asJava {
		m.WriteJava()
	} else {
		m.WriteGo()
	}
}
