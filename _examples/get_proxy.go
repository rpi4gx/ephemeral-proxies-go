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

	// Ensures user has provided their RapidAPI key
	if len(*apiKey) == 0 {
		usage()
		os.Exit(-1)
	}

	// Obtains a proxy available for the next 30 mins
	proxy, err := ephemeralproxies.NewProxy(*apiKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Details of new proxy obtained:")
	fmt.Println(proxy)
	fmt.Println()
}
