package main

import (
	"flag"
	"fmt"
	"os"

	ephemeralproxies "github.com/rpi4gx/ephemeral-proxies-go"
)

func usage() {
	fmt.Println("required arguments")
	fmt.Println("\t--key=RAPIDAPI_KEY")
}
func main() {

	var apiKey = flag.String("key", "", "User's RapidAPI Key")
	flag.Parse()

	if len(*apiKey) == 0 {
		usage()
		os.Exit(-1)
	}

	p, err := ephemeralproxies.NewProxy(*apiKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
