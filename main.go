package main

import (
	"log"
	"os"
)

const (
	geoIpDBFileName   = "GeoLite2-City.mmdb"
	addressesFileName = "addresses.txt"
	resultsFileName   = "results.csv"
)

func main() {
	addresses, err := readFile(addressesFileName)
	if err != nil {
		log.Fatal(err)
	}
	results := geoip(geoIpDBFileName, addresses)

	if len(results) == 0 {
		log.Println("No results found")
		os.Exit(0)
	}

	if len(results) != len(addresses) {
		log.Printf("Expected %d results, got %d\n", len(addresses), len(results))
	}

	if err = writeFile(resultsFileName, results); err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
