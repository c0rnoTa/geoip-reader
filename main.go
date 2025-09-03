package main

import (
	"flag"
	"log"
	"os"
)

var (
	Version   = "master"
	BuildTime = "unknown"
	Commit    = "unknown"
)

const (
	geoIpDBFileName   = "GeoLite2-City.mmdb"
	addressesFileName = "addresses.txt"
	resultsFileName   = "results.csv"
)

func main() {

	log.Printf("GeoIP reader version: %s (%s build at %s)", Version, Commit, BuildTime)

	// Определяем флаги
	geoDb := flag.String("d", geoIpDBFileName, "MMDB database file")
	addressesFile := flag.String("i", addressesFileName, "Input addresses file")
	resultsFile := flag.String("o", resultsFileName, "Output results file")

	// Парсим аргументы
	flag.Parse()

	addresses, err := readFile(*addressesFile)
	if err != nil {
		log.Fatal(err)
	}
	results := geoip(*geoDb, addresses)

	if len(results) == 0 {
		log.Println("No results found")
		os.Exit(0)
	}

	if len(results) != len(addresses) {
		log.Printf("Expected %d results, got %d\n", len(addresses), len(results))
	}

	if err = writeFile(*resultsFile, results); err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
