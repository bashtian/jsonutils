package jsonutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var Writer io.Writer = os.Stdout

func Get(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return b, err
}

func ParseJson(b []byte) (interface{}, error) {
	var f interface{}
	err := json.Unmarshal(b, &f)
	return f, err
}

func PrintGo(f interface{}) {
	fu := func(m map[string]interface{}) { parseMap(m) }
	print(f, fu, "type %s []struct {\n", "type %s struct {\n")
}

func PrintJava(f interface{}) {
	fu := func(m map[string]interface{}) {
		v, n := parseMapJava(m)
		if v != nil {
			parseArrayJava(v, n)
		}
	}
	fmt.Fprintln(Writer, "import com.google.gson.annotations.SerializedName;\n")
	print(f, fu, "//NOTE: use as an array\nclass %s {\n", "class %s {\n")
}

func print(f interface{}, fu func(map[string]interface{}), array string, object string) {
	var m map[string]interface{}
	switch v := f.(type) {
	case []interface{}:
		m = v[0].(map[string]interface{})
		fmt.Fprintf(Writer, array, "Data")
	default:
		m = f.(map[string]interface{})
		fmt.Fprintf(Writer, object, "Data")
	}
	fu(m)
	fmt.Fprintln(Writer, "}")
}

func parseMap(m map[string]interface{}) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			printType(k, "string")
		case bool:
			printType(k, "bool")
		case float64:
			if float64(int64(vv)) == vv {
				printType(k, "int64")
			} else {
				printType(k, "float64")
			}
		case []interface{}:
			if len(vv) > 0 {
				switch vvv := vv[0].(type) {
				case string:
					printType(k, "[]string")
				case []interface{}:
					printObject(k, "[]struct", func() { parseMap(vvv[0].(map[string]interface{})) })
				case map[string]interface{}:
					printObject(k, "[]struct", func() { parseMap(vvv) })
				default:
					printType(k, "interface{}")
				}
			} else {
				// empty array
				printType(k, "[]interface{}")
			}
		case map[string]interface{}:
			printObject(k, "struct", func() { parseMap(vv) })
		default:
			printType(k, "interface{}")
		}
	}
}

func parseMapJava(m map[string]interface{}) ([]map[string]interface{}, []string) {
	var data []map[string]interface{}
	var names []string
	for k, v := range m {
		name := replaceName(k)
		switch vv := v.(type) {
		case string:
			printValuesJava("String", k)
		case float64:
			if float64(int(vv)) == vv {
				printValuesJava("int", k)
			} else {
				printValuesJava("double", k)
			}
		case bool:
			printValuesJava("boolean", k)
		case []interface{}:
			printValuesJava(name+"[]", k)
			data = append(data, vv[0].(map[string]interface{}))
			names = append(names, k)
		case map[string]interface{}:
			printValuesJava(name, k)
			data = append(data, vv)
			names = append(names, k)
		default:
			printValuesJava("Object", k)
		}
	}
	return data, names
}

func printType(n string, t string) {
	name := replaceName(n)
	fmt.Fprintf(Writer, "%s %s `json:\"%s\"`\n", name, t, n)
}

func printObject(n string, t string, f func()) {
	name := replaceName(n)
	fmt.Fprintf(Writer, "%s %s {\n", name, t)
	f()
	fmt.Fprintf(Writer, "} `json:\"%s\"`\n", n)
}

func parseArrayJava(m []map[string]interface{}, s []string) {
	for i, v := range m {
		fmt.Fprintln(Writer, "class", s[i], "{")
		v, n := parseMapJava(v)
		fmt.Fprintln(Writer, "}")
		if v != nil {
			parseArrayJava(v, n)
		}
	}
}

func printValuesJava(javaType, key string) {
	const tmpl = `
@SerializedName("{{.Key}}")
private {{.Type}} {{.LowerName}};

public {{.Type}} get{{.Name}}() {
	return {{.LowerName}};
}

public void set{{.Name}}({{.Type}} {{.LowerName}}) {
	this.{{.LowerName}} = {{.LowerName}};
}
`
	tmpName := replaceName(key)
	data := struct {
		Type      string
		Key       string
		Name      string
		LowerName string
	}{
		javaType,
		key,
		tmpName,
		strings.ToLower(tmpName),
	}
	t := template.Must(template.New("type").Parse(tmpl))
	t.Execute(os.Stdout, data)
}

func replaceName(n string) string {
	for _, c := range "@_-+.,!" {
		n = strings.Replace(n, string(c), " ", -1)
	}
	n = strings.Title(n)
	n = strings.Replace(n, " ", "", -1)
	return n
}

func Mock(b []byte, i interface{}) ([]byte, error) {
	err := json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package main\nvar test = %#v", i)

	form, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	//err = ioutil.WriteFile("test_mock.go", buf.Bytes(), 0644)
	//if err != nil {
	//	return nil, err
	//}

	return form, nil
}
