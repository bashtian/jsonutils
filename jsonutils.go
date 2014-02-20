package jsonutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Model struct {
	Writer      io.Writer
	WithExample bool
	Data        interface{}
	Name        string
	Format      bool
	Convert     bool
}

func New(data interface{}, name string) *Model {
	if name == "" {
		name = "Data"
	}
	return &Model{
		Writer:      os.Stdout,
		WithExample: true,
		Data:        data,
		Name:        name,
		Format:      true,
		Convert:     true,
	}
}

func FromBytes(bytes []byte, name string) (*Model, error) {
	f, err := ParseJson(bytes)
	if err != nil {
		return nil, err
	}
	return New(f, name), nil
}

func GetModel(url string) (*Model, error) {
	b, name, err := Get(url)
	if err != nil {
		return nil, err
	}
	return FromBytes(b, name)
}

func Get(url string) ([]byte, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Add("Accept", "application/json")
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, "", err
	}
	return b, getName(url), err
}

func ParseJson(b []byte) (interface{}, error) {
	var f interface{}
	err := json.Unmarshal(b, &f)
	return f, err
}

func PrintGo(f interface{}, name string) {
	WriteGo(os.Stdout, f, name)
}

func WriteGo(w io.Writer, f interface{}, name string) {
	m := &Model{
		Writer:      w,
		WithExample: true,
		Data:        f,
		Name:        name,
	}
	m.WriteGo()
}

func (m *Model) WriteGo() {
	if m.Format {
		org := m.Writer

		var buf bytes.Buffer
		m.Writer = &buf

		m.writeGo()

		b, err := format.Source(buf.Bytes())
		if err == nil {
			org.Write(b)
		} else {
			io.Copy(org, &buf)
		}
		m.Writer = org
	} else {
		m.writeGo()
	}

}

func (m *Model) writeGo() {
	fu := func(ms map[string]interface{}) { m.parseMap(ms) }
	m.print(fu, "type %s []struct {\n", "type %s struct {\n")
}

func (m *Model) WriteJava() {
	fu := func(ms map[string]interface{}) {
		v, n := m.parseMapJava(ms)
		if v != nil {
			m.parseArrayJava(v, n)
		}
	}
	fmt.Fprintln(m.Writer, "import com.google.gson.annotations.SerializedName;\n")
	m.print(fu, "//NOTE: use as an array\npublic class %s {\n", "class %s {\n")
}

func (m *Model) print(fu func(map[string]interface{}), array, object string) {
	var ma map[string]interface{}
	switch v := m.Data.(type) {
	case []interface{}:
		ma = v[0].(map[string]interface{})
		fmt.Fprintf(m.Writer, array, m.Name)
	default:
		ma = m.Data.(map[string]interface{})
		fmt.Fprintf(m.Writer, object, m.Name)
	}
	fu(ma)
	fmt.Fprintln(m.Writer, "}")
}

func (m *Model) parseMap(ms map[string]interface{}) {
	keys := getSortedKeys(ms)
	for _, k := range keys {
		switch vv := ms[k].(type) {
		case string:
			if m.Convert {
				t, converted := parseType(vv)
				m.printType(k, vv, t, converted)
			} else {
				m.printType(k, vv, "string", false)
			}
		case bool:
			m.printType(k, vv, "bool", false)
		case float64:
			//json parser always returns a float for number values, check if it is an int value
			if float64(int64(vv)) == vv {
				m.printType(k, vv, "int64", false)
			} else {
				m.printType(k, vv, "float64", false)
			}
		case int64:
			m.printType(k, vv, "int64", false)
		case []interface{}:
			if len(vv) > 0 {
				switch vvv := vv[0].(type) {
				case string:
					m.printType(k, vv[0], "[]string", false)
				case float64:
					m.printType(k, vv[0], "[]float64", false)
				case []interface{}:
					m.printObject(k, "[]struct", func() { m.parseMap(vvv[0].(map[string]interface{})) })
				case map[string]interface{}:
					m.printObject(k, "[]struct", func() { m.parseMap(vvv) })
				default:
					//fmt.Printf("unknown type: %T", vvv)
					m.printType(k, nil, "interface{}", false)
				}
			} else {
				// empty array
				m.printType(k, nil, "[]interface{}", false)
			}
		case map[string]interface{}:
			m.printObject(k, "struct", func() { m.parseMap(vv) })
		default:
			//fmt.Printf("unknown type: %T", vv)
			m.printType(k, nil, "interface{}", false)
		}
	}
}

func (m *Model) parseMapJava(ms map[string]interface{}) ([]map[string]interface{}, []string) {
	var data []map[string]interface{}
	var names []string
	for k, v := range ms {
		name := replaceName(k)
		switch vv := v.(type) {
		case string:
			m.printValuesJava("String", k)
		case float64:
			if float64(int(vv)) == vv {
				m.printValuesJava("int", k)
			} else {
				m.printValuesJava("double", k)
			}
		case bool:
			m.printValuesJava("boolean", k)
		case []interface{}:
			if len(vv) > 0 {
				switch vvv := vv[0].(type) {
				case string:
					m.printValuesJava("String[]", k)
				case float64:
					m.printValuesJava("float[]", k)
				case []interface{}:
					m.printValuesJava(name+"[]", k)
					data = append(data, vvv[0].(map[string]interface{}))
					names = append(names, k)
				case map[string]interface{}:
					m.printValuesJava(name+"[]", k)
					data = append(data, vvv)
					names = append(names, k)
				default:
					//fmt.Printf("unknown type: %T", vvv)
					m.printValuesJava("Object", k)
				}
			} else {
				// empty array
				m.printValuesJava("Object[]", k)
			}

		case map[string]interface{}:
			m.printValuesJava(name, k)
			data = append(data, vv)
			names = append(names, k)
		default:
			m.printValuesJava("Object", k)
		}
	}
	return data, names
}

func parseType(value string) (string, bool) {
	if _, err := time.Parse(time.RFC3339, value); err == nil {
		return "time.Time", false
	} else if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return "int64", true
	} else if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "float64", true
	} else if _, err := strconv.ParseBool(value); err == nil {
		return "bool", true
	} else {
		return "string", false
	}
}

func (m *Model) printType(key string, value interface{}, t string, converted bool) {
	name := replaceName(key)
	if converted {
		key += ",string"
	}
	if m.WithExample {
		fmt.Fprintf(m.Writer, "%s %s `json:\"%s\"` // %v\n", name, t, key, value)
	} else {
		fmt.Fprintf(m.Writer, "%s %s `json:\"%s\"`\n", name, t, key)
	}
}

func (m *Model) printObject(n string, t string, f func()) {
	name := replaceName(n)
	fmt.Fprintf(m.Writer, "%s %s {\n", name, t)
	f()
	fmt.Fprintf(m.Writer, "} `json:\"%s\"`\n", n)
}

func (m *Model) parseArrayJava(ms []map[string]interface{}, s []string) {
	for i, v := range ms {
		fmt.Fprintln(m.Writer, "public class", replaceName(s[i]), "{")
		v, n := m.parseMapJava(v)
		fmt.Fprintln(m.Writer, "}")
		if v != nil {
			m.parseArrayJava(v, n)
		}
	}
}

func (m *Model) printValuesJava(javaType, key string) {
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
	t.Execute(m.Writer, data)
}

func replaceName(n string) string {
	for _, c := range "@#_-+.,!$:/\\" {
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

func getSortedKeys(m map[string]interface{}) (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return
}

func getName(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return "Data"
	}
	s := strings.Split(p.Path, "/")
	if len(s) < 1 {
		return "Data"
	}
	return strings.Title(s[len(s)-1])
}
