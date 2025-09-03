# Check IP address list in GeoIP

## Args

* **d** - MMDB database file (default "GeoLite2-City.mmdb")
* **i** - Input addresses file (default "addresses.txt")
* **o** - Output results file (default "results.csv")

## Usage

1. Create `addresses.txt` with one column of IPs
2. Place `GeoLite2-City.mmdb` on the same dir form https://cdn.jsdelivr.net/npm/geolite2-city/ or https://db-ip.com/db/download/ip-to-country-lite
3. Run this app. It will create `result.csv` with network and iso code of country.