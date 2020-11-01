package triangulate

import (
	"geoipd/pkg/libgeoip"
	"geoipd/pkg/libgeoip/filters"
	"net"

	"github.com/monzo/typhon"
)

func PostTriangulateHandler(app libgeoip.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		rawRequest := req.Context.Value(filters.ValidationResult)
		tReq := rawRequest.(*PostTriangulateRequest)
		m := map[string]Location{}

		for _, ip := range tReq.IPs {
			city, err := app.DB.City(net.IP(ip))
			if err != nil {
				continue
			}
			m[net.IP.String(net.IP(ip))] = Location{
				City:      city.City.Names["en"],
				Longitude: city.Location.Longitude,
				Latitude:  city.Location.Latitude,
				AccuracyRadius: city.Location.AccuracyRadius,
			}
		}

		response := req.Response(&PostTriangulateResponse{
			IPs: m,
		})

		response.StatusCode = 200
		return response
	}
}
