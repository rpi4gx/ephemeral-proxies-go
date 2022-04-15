[![Go Reference](https://pkg.go.dev/badge/github.com/rpi4gx/ephemeral-proxies-go.svg)](https://pkg.go.dev/github.com/rpi4gx/ephemeral-proxies-go)
## Golang client library for [Ephemeral proxies API](https://www.ephemeral-proxies.net/)

:warning: This library requires a valid Rapid API key to access Ephemeral Proxies API. A Rapid API key can easily be obtained on https://rapidapi.com/.

:information_source: More information about Ephemeral Proxies API can be found [here](https://rapidapi.com/rpi4gx/api/ephemeral-proxies)

### Quick start

```
go get github.com/rpi4gx/ephemeral-proxies-go
```

myapp.go:
```
package main

import (
	"fmt"
	proxy "github.com/rpi4gx/ephemeral-proxies-go"
)

func main() {
	p, err := proxy.NewProxy("REPLACE_WITH_RAPIDAPI_KEY)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
```

### Running the examples

Clone the repo

```
$ git clone https://github.com/rpi4gx/ephemeral-proxies-go.git
```
Example 1: Pull a new proxy from the API
```
$ go run _examples/get_proxy.go --key=RAPIDAPI_KEY
Details of new proxy obtained:
{
    "id": "3dd035fb5cee360fb0dc3c8cdba927fc",
    "host": "lc13e.ep-proxy.net",
    "port": 32793,
    "expires_at": "2022-04-15T20:24:40Z",
    "whitelisted_ips": [
        "219.182.31.4"
    ],
    "visibility": {
        "ip": "31.40.246.1",
        "country": "United Kingdom",
        "country_iso": "GB",
        "country_eu": true,
        "latitude": 51.5164,
        "longitude": -0.093,
        "timezone": "Europe/London",
        "asn": "",
        "asn_org": "",
        "zip_code": "EC2V",
        "region_name": "England",
        "region_code": "ENG",
        "city": "London"
    },
    "features": {
        "static": true,
        "supported_protocols": {
            "socks4": false,
            "socks5": false,
            "http": true,
            "https": true
        }
    }
}
```
Example 2: Get API service status
```
$ go run _examples/service_status.go --key=RAPIDAPI_KEY
Ephemeral Proxies Service Status:
{
    "availability": {
        "total": {
            "proxies": 822
        },
        "by_country": [
            {
                "country_iso": "BE",
                "proxies": 4
            },
            {
                "country_iso": "AT",
                "proxies": 12
            },
            {
                "country_iso": "PL",
                "proxies": 12
            },
            {
                "country_iso": "GB",
                "proxies": 286
            },
            {
                "country_iso": "NL",
                "proxies": 22
            },
            {
                "country_iso": "SK",
                "proxies": 9
            },
            {
                "country_iso": "IS",
                "proxies": 7
            },
            {
                "country_iso": "ES",
                "proxies": 6
            },
            {
                "country_iso": "RU",
                "proxies": 207
            },
            {
                "country_iso": "BR",
                "proxies": 113
            },
            {
                "country_iso": "EE",
                "proxies": 2
            },
            {
                "country_iso": "SA",
                "proxies": 14
            },
            {
                "country_iso": "DE",
                "proxies": 118
            },
            {
                "country_iso": "FR",
                "proxies": 10
            }
        ]
    }
}
```


