package main

import (
	"github.com/ahfuzhang/maxminddb-golang"
	"log"
	"net"
)

func geoip(geoipDb string, addresses []string) map[string]string {
	db, err := maxminddb.Open(geoipDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results := make(map[string]string)

	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	}

	for _, address := range addresses {

		ip := net.ParseIP(address)
		network, _, err := db.LookupNetwork(ip, &record)
		if err != nil {
			log.Printf("unable to process address '%s': %v", address, err)
			continue
		}
		results[network.String()] = record.Country.ISOCode
	}

	return results
}
