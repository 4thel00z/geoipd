package triangulate

import (
	"geoipd/pkg/libgeoip"
	"net"
)

func Locations(app libgeoip.App, ips... IP) map[string]Location {
	m := map[string]Location{}

	for _, ip := range ips {
		city, err := app.DB.City(net.IP(ip))
		if err != nil {
			continue
		}
		m[net.IP.String(net.IP(ip))] = Location{
			City:           city.City.Names["en"],
			Longitude:      city.Location.Longitude,
			Latitude:       city.Location.Latitude,
			AccuracyRadius: city.Location.AccuracyRadius,
		}
	}
	return m
}
