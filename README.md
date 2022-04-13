## Go package to make use of [Ephemeral proxies API](https://www.ephemeral-proxies.net/)


Example

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


