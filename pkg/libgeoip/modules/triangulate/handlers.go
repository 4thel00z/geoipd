package triangulate

import (
	"encoding/json"
	"fmt"
	"geoipd/pkg/libgeoip"
	"geoipd/pkg/libgeoip/filters"
	"github.com/fapian/geojson2svg/pkg/geojson2svg"
	"github.com/monzo/typhon"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmapi"
	"github.com/paulmach/osm/osmgeojson"
	"net"
	"net/http"
	"os"
)

func PostLocateHandler(app libgeoip.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		rawRequest := req.Context.Value(filters.ValidationResult)
		tReq := rawRequest.(*PostLocateRequest)
		response := req.Response(&PostLocateResponse{
			IPs: Locations(app, tReq.IPs...),
		})

		response.StatusCode = 200
		return response
	}
}

func PostRenderHandler(app libgeoip.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		rawRequest := req.Context.Value(filters.ValidationResult)
		tReq := rawRequest.(*PostRenderRequest)
		l := Locations(app, tReq.IP)
		location := l[net.IP.String(net.IP(tReq.IP))]

		delta := tReq.Padding

		lat, lon := location.Latitude, location.Longitude
		bounds := &osm.Bounds{
			MinLat: lat - delta, MaxLat: lat + delta,
			MinLon: lon - delta, MaxLon: lon + delta,
		}

		o, err := osmapi.Map(req.Context, bounds) // fetch data from the osm api.
		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}

		geo, err := osmgeojson.Convert(o,
			osmgeojson.IncludeInvalidPolygons(false),
			osmgeojson.NoMeta(true),
		)

		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}

		for _, feature := range geo.Features {
			err := Styles.AddToFeature(feature)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s", err)
				continue
			}
		}
		gj, err := json.Marshal(geo)

		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}

		svg := geojson2svg.New()
		err = svg.AddFeatureCollection(string(gj))
		fmt.Println(string(gj))

		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}

		point := geojson.NewFeature(orb.Point{
			lat, lon,
		})

		err = Styles.AddToFeature(point)

		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}

		gj, err = json.Marshal(point)
		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}
		err = svg.AddFeatureCollection(string(gj))

		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}
		response := req.Response(nil)

		_, err = response.Write([]byte(svg.Draw(tReq.Width, tReq.Height, geojson2svg.UseProperties(Styles.Keys()))))
		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}
		response.Header.Add("Content-Type", "image/svg+xml")
		response.StatusCode = 200
		return response
	}
}
