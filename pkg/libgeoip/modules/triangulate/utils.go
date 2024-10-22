package triangulate

import (
	"errors"
	"geoipd/pkg/libgeoip"
	"image/color"
	"net"
)

func Locations(app libgeoip.App, ips ...IP) map[string]Location {
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

var errInvalidFormat = errors.New("invalid format")

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
