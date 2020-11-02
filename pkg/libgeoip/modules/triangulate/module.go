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
		"locate": {
			Path:        "locate",
			Method:      "POST",
			CurlExample: "http POST http://<addr>/<version>/<namespace>/<path> < examples/ips.json",
			Service:     PostLocateHandler,
			Validator:   libgeoip.GenerateRequestValidator(PostLocateRequest{}),
		},
		"render": {
			Path:        "render",
			Method:      "POST",
			CurlExample: "http POST http://<addr>/<version>/<namespace>/<path> < examples/render.json > out.svg",
			Service:     PostRenderHandler,
			Validator:   libgeoip.GenerateRequestValidator(PostRenderRequest{}),
		},
	}
}

func (t Triangulate) LongPath(route libgeoip.Route) string {
	return libgeoip.DefaultLongPath(t, route)
}
