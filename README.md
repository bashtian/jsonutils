jsonutils
=========

Converter for JSON data to a Go struct or a Java class for GSON

	go get github.com/bashtian/jsonutils/jsonutils

You can print the structure of a JSON from a URL

	jsonutils https://api.github.com/users/bashtian/repos

or from a file

	jsonutils -f file.json

If you want to print example data as comments, use the -x parmeter

	jsonutils -x https://api.github.com/users/bashtian/repos

You can also print Java code 

	jsonutils -j https://api.github.com/users/bashtian/repos

###Example
####JSON
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
####Go
	
	jsonutils -x -f example.json

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
			LastCall time.Time `json:"last_call"` // 2013-01-10T05:27:07Z
			Number   string    `json:"number"`    // 212 555-1234
			Type     string    `json:"type"`      // home
		} `json:"phoneNumber"`
		Tags []string `json:"tags"` // music
	}
