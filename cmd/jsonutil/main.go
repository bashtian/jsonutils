package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v1"

	"github.com/bashtian/jsonutils"
)

var (
	filename      = flag.String("f", "", "use a file as input")
	withExamples  = flag.Bool("x", false, "print examples as comment")
	convertString = flag.Bool("c", true, "convert strings (go only)")
	asJava        = flag.Bool("j", false, "print Java instead of Go code")
	isYaml        = flag.Bool("y", false, "input is YAML")
)

func main() {
	flag.Parse()
	url := flag.Arg(0)

	var m *jsonutils.Model
	var err error
	var data []byte
	var name string

	if *filename != "" {
		name = strings.Split(*filename, ".")[0]
		data, err = ioutil.ReadFile(*filename)
	} else if url != "" {
		data, name, err = jsonutils.Get(url)
	} else {
		data, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatal(err)
	}

	if !*isYaml {
		m, err = jsonutils.FromBytes(data, strings.Title(name))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var f map[string]interface{}
		err = yaml.Unmarshal(data, &f)
		if err != nil {
			log.Fatal("goyaml.Unmarshal:", err)
		}
		m = jsonutils.New(f, strings.Title(name))
	}

	m.WithExample = *withExamples
	m.Convert = *convertString
	if *asJava {
		m.WriteJava()
	} else {
		m.WriteGo()
	}
}
