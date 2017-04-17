jsonutils
=========

Converter for JSON data to a Go struct or a Java class for GSON

```bash
go get github.com/bashtian/jsonutils/cmd/jsonutil
```

You can print the structure of a JSON from a URL

```bash
jsonutil https://api.github.com/repos/bashtian/jsonutils
```

or from a file

```bash
jsonutil -f file.json
```

or from stdin

```bash
echo '{"some": "data"}' | jsonutil
```

If you want to print example data as comments, use the -x parameter

```bash
jsonutil -x https://api.github.com/repos/bashtian/jsonutils
```

You can also print Java code 

```bash
jsonutil -j https://api.github.com/repos/bashtian/jsonutils
```

### Example
#### JSON
```json
{
    "firstName": "John",
    "lastName": "Smith",
    "age": 25,
    "balance": 123.45,
    "address": {
        "streetAddress": "21 2nd Street",
        "city": "New York",
        "state": "NY",
        "postalCode": "10021"
    },
    "phoneNumber": [
        {
            "type": "home",
            "number": "212 555-1234",
            "last_call": "2013-01-10T05:27:07Z"
        },
        {
            "type": "fax",
            "number": "646 555-4567",
            "last_call": "2013-01-10T05:27:07Z"
        }
    ],
    "tags": ["music","video"]
}
```
#### Go
	
	jsonutil -x -c=false -f Example.json
	
```go
type Example struct {
	Address struct {
		City          string `json:"city"`          // New York
		PostalCode    string `json:"postalCode"`    // 10021
		State         string `json:"state"`         // NY
		StreetAddress string `json:"streetAddress"` // 21 2nd Street
	} `json:"address"`
	Age         int64   `json:"age"`       // 25
	Balance     float64 `json:"balance"`   // 123.45
	FirstName   string  `json:"firstName"` // John
	LastName    string  `json:"lastName"`  // Smith
	PhoneNumber []struct {
		LastCall string `json:"last_call"` // 2013-01-10T05:27:07Z
		Number   string `json:"number"`    // 212 555-1234
		Type     string `json:"type"`      // home
	} `json:"phoneNumber"`
	Tags []string `json:"tags"` // music
}
```

	
	jsonutils -f Example.json
	
```go
type Example struct {
	Address struct {
		City          string `json:"city"`
		PostalCode    int64  `json:"postalCode,string"`
		State         string `json:"state"`
		StreetAddress string `json:"streetAddress"`
	} `json:"address"`
	Age         int64   `json:"age"`
	Balance     float64 `json:"balance"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	PhoneNumber []struct {
		LastCall time.Time `json:"last_call"`
		Number   string    `json:"number"`
		Type     string    `json:"type"`
	} `json:"phoneNumber"`
	Tags []string `json:"tags"`
}

```
