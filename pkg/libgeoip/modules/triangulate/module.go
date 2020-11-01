package triangulate

import (
	"geoipd/pkg/libgeoip"
)

type Triangulate struct{}

var (
	Module = Triangulate{}
)

func (t Triangulate) Version() string {
	return "v1"
}

func (t Triangulate) Namespace() string {
	return "triangulate"
}

func (t Triangulate) Routes() map[string]libgeoip.Route {
	// Add route definitions here
	return map[string]libgeoip.Route{
		"location": {
			Path:        "location",
			Method:      "POST",
			CurlExample: "http POST http://<addr>/<version>/<namespace>/<path> < examples/ips.txt",
			Service:     PostTriangulateHandler,
			Validator:   libgeoip.GenerateRequestValidator(PostTriangulateRequest{}),
		},
	}
}

func (t Triangulate) LongPath(route libgeoip.Route) string {
	return libgeoip.DefaultLongPath(t, route)
}
