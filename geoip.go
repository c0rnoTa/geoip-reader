package main

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func geoip(geoipDb string, addresses []string) map[string]*geoip2.City {
	db, err := geoip2.Open(geoipDb)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	results := make(map[string]*geoip2.City)

	for _, address := range addresses {
		ip := net.ParseIP(address)
		record, err := db.City(ip)
		if err != nil {
			log.Printf("unable to process address '%s': %v", address, err)
			continue
		}
		results[address] = record
	}
	return results
}
