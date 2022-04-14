[![Go Reference](https://pkg.go.dev/badge/github.com/rpi4gx/ephemeral-proxies-go.svg)](https://pkg.go.dev/github.com/rpi4gx/ephemeral-proxies-go)
## Golang client library for [Ephemeral proxies API](https://www.ephemeral-proxies.net/)

:warning: This library requires a valid Rapid API key to access Ephemeral Proxies API. A Rapid API key can easily be obtained on https://rapidapi.com/.

:information_source: More information about Ephemeral Proxies API can be found [here](https://rapidapi.com/rpi4gx/api/ephemeral-proxies)

### Quick start

```
go mod init sample
go get github.com/rpi4gx/ephemeral-proxies-go
```

sample.go:
```
package main

import (
	"fmt"

	proxy "github.com/rpi4gx/ephemeral-proxies-go"
)

func main() {
	p, err := proxy.NewProxy("PASTE_YOUR_RAPIDAPI_KEY_HERE")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
```
// TODO
```
go run sample.go

```


