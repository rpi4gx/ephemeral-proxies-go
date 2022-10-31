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

	// Ensure user has provided their RapidAPI key
	if len(*apiKey) == 0 {
		usage()
		os.Exit(-1)
	}

	// Obtains information about the current state of the Ephemeral Proxies API service
	balance, err := ephemeralproxies.GetUserBalance(*apiKey, ephemeralproxies.Residential)
	if err != nil {
		panic(err)
	}
	fmt.Println("User's balance:")
	fmt.Println(balance)
}
