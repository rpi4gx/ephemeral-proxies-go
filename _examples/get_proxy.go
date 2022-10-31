package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	ephemeralproxies "github.com/rpi4gx/ephemeral-proxies-go"
)

func usage() {
	fmt.Println("required arguments")
	fmt.Println("\t--key=RAPIDAPI_KEY")
	fmt.Println("optional arguments")
	fmt.Println("\t--proxy_type=<datacenter|residential>")
}

func getProxyType(p string) (ephemeralproxies.ProxyType, error) {
	if p == "residential" {
		return ephemeralproxies.Residential, nil
	} else if p == "datacenter" {
		return ephemeralproxies.Datacenter, nil
	}
	return 0, errors.New("invalid proxy type")
}

func main() {

	var apiKey = flag.String("key", "", "User's RapidAPI Key")
	var pType = flag.String("type", "datacenter", "Type of proxy, 'datacenter' or 'residential'")
	flag.Parse()

	// Ensures user has provided their RapidAPI key
	if len(*apiKey) == 0 {
		usage()
		os.Exit(-1)
	}

	// If proxy type was set, check that it is a valid type
	var proxyType, err = getProxyType(*pType)
	if err != nil {
		usage()
		os.Exit(-1)
	}

	// Obtains a proxy available for the next 30 mins
	proxy, err := ephemeralproxies.NewProxy(*apiKey, proxyType)
	if err != nil {
		panic(err)
	}
	fmt.Println("Details of new proxy obtained:")
	fmt.Println(proxy)
	fmt.Println()
}
