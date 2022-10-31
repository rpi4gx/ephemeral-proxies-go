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
	p, err := proxy.NewProxy("REPLACE_WITH_RAPIDAPI_KEY, ephemeralproxies.Residential)
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
$ cd ephemeral-proxies-go
```
Example 1: Pull a new proxy from the API
```
$ go run _examples/get_proxy.go --key=RAPIDAPI_KEY --type=datacenter
Details of new proxy obtained:
{
    "id": "61b6f0db6c79077c539aca49e9001811",
    "host": "l56ff.ep-proxy.net",
    "port": 40802,
    "expires_at": "2022-10-31T10:04:34Z",
    "whitelisted_ips": [
        "18.19.280.40"
    ],
    "visibility": {
        "ip": "185.102.113.55",
        "country": "Brazil",
        "country_iso": "BR",
        "country_eu": false,
        "latitude": -22.9072,
        "longitude": -43.1883,
        "timezone": "America/Sao_Paulo",
        "asn": "AS35830",
        "asn_org": "BTT Group Finance Ltd",
        "zip_code": "",
        "region_name": "Rio de Janeiro",
        "region_code": "RJ",
        "city": "Rio de Janeiro"
    },
    "features": {
        "static": true,
        "type": "datacenter",
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
$ go run _examples/service_status.go --key=RAPIDAPI_KEY --type=residential
Ephemeral Proxies Service Status:
{
    "availability": {
        "total": {
            "proxies": 4480591
        },
        "by_country": [
            {
                "country_iso": "US",
                "proxies": 139547
            },
            {
                "country_iso": "NL",
                "proxies": 140091
            },
            {
                "country_iso": "IN",
                "proxies": 140080
            },
            {
                "country_iso": "JP",
                "proxies": 139903
            },
            {
                "country_iso": "MX",
                "proxies": 139948
            },
            {
                "country_iso": "PL",
                "proxies": 140327
            },
            {
                "country_iso": "PK",
                "proxies": 139980
            },
            {
                "country_iso": "IT",
                "proxies": 139796
            },
            {
                "country_iso": "DE",
                "proxies": 140017
            },
            {
                "country_iso": "FR",
                "proxies": 140231
            },
            {
                "country_iso": "PH",
                "proxies": 140056
            },
            {
                "country_iso": "SE",
                "proxies": 140106
            },
            {
                "country_iso": "TW",
                "proxies": 140245
            },
            {
                "country_iso": "RU",
                "proxies": 140256
            },
            {
                "country_iso": "GB",
                "proxies": 139952
            },
            {
                "country_iso": "BR",
                "proxies": 139616
            },
            {
                "country_iso": "AU",
                "proxies": 140454
            },
            {
                "country_iso": "CA",
                "proxies": 140076
            },
            {
                "country_iso": "HK",
                "proxies": 139827
            },
            {
                "country_iso": "PT",
                "proxies": 139564
            },
            {
                "country_iso": "BE",
                "proxies": 140379
            },
            {
                "country_iso": "TR",
                "proxies": 139992
            },
            {
                "country_iso": "KR",
                "proxies": 140288
            },
            {
                "country_iso": "ES",
                "proxies": 139684
            },
            {
                "country_iso": "SG",
                "proxies": 140350
            },
            {
                "country_iso": "MY",
                "proxies": 139817
            },
            {
                "country_iso": "CH",
                "proxies": 139966
            },
            {
                "country_iso": "AR",
                "proxies": 140282
            },
            {
                "country_iso": "TH",
                "proxies": 140169
            },
            {
                "country_iso": "UA",
                "proxies": 139917
            },
            {
                "country_iso": "GR",
                "proxies": 139654
            },
            {
                "country_iso": "IL",
                "proxies": 140021
            }
        ]
    }
}
```
Example 3: Get monthly user's balance
```
$ go run _examples/user_balance.go --key=RAPIDAPI_KEY
User's balance:
{
    "consumed_megabytes": 845,
    "limit_megabytes": 50000
}
