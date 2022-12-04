package traefikgeoip2

import (
	"fmt"
	"net"

	"github.com/IncSW/geoip2"
)

// Unknown constant for undefined data.
const Unknown = "XX"

// DefaultDBPath default GeoIP2 database path.
const DefaultDBPath = "GeoLite2-Country.mmdb"

const (
	// CountryHeader country header name.
	CountryHeader = "X-GeoIP2-Country"
	// RegionHeader region header name.
	RegionHeader = "X-GeoIP2-Region"
	// CityHeader city header name.
	CityHeader = "X-GeoIP2-City"
	// LongitudeHeader longitude position of request.
	LongitudeHeader = "X-GeoIP2-Long"
	// LatitudeHeader Latitude position of request.
	LatitudeHeader = "X-GeoIP2-Lat"
)

// GeoIPResult GeoIPResult.
type GeoIPResult struct {
	country   string
	region    string
	city      string
	latitude  float64
	longitude float64
}

// LookupGeoIP2 LookupGeoIP2.
type LookupGeoIP2 func(ip net.IP) (*GeoIPResult, error)

// CreateCityDBLookup CreateCityDBLookup.
func CreateCityDBLookup(rdr *geoip2.CityReader) LookupGeoIP2 {
	return func(ip net.IP) (*GeoIPResult, error) {
		rec, err := rdr.Lookup(ip)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		retval := GeoIPResult{
			country:   rec.Country.ISOCode,
			region:    Unknown,
			city:      Unknown,
			latitude:  rec.Location.Latitude,
			longitude: rec.Location.Longitude,
		}
		if city, ok := rec.City.Names["en"]; ok {
			retval.city = city
		}
		if rec.Subdivisions != nil {
			retval.region = rec.Subdivisions[0].ISOCode
		}
		return &retval, nil
	}
}

// CreateCountryDBLookup CreateCountryDBLookup.
func CreateCountryDBLookup(rdr *geoip2.CountryReader) LookupGeoIP2 {
	return func(ip net.IP) (*GeoIPResult, error) {
		rec, err := rdr.Lookup(ip)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		retval := GeoIPResult{
			country: rec.Country.ISOCode,
			region:  Unknown,
			city:    Unknown,
		}
		return &retval, nil
	}
}
