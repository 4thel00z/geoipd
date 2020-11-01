package debug

import (
	"geoipd/pkg/libgeoip"
	"geoipd/pkg/libgeoip/jwt"
)

type Debug struct{}

var (
	Module = Debug{}
)

func (Y Debug) Version() string {
	return "v1"
}

func (Y Debug) Namespace() string {
	return "debug"
}

func (Y Debug) Routes() map[string]libgeoip.Route {
	// Add route definitions here
	return map[string]libgeoip.Route{
		"routes": {
			Path:        "routes",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/routes",
			Service:     GetRoutesHandler,
		},
		"private": {
			Path:        "private",
			Method:      "GET",
			CurlExample: "curl http://<addr>/<version>/<namespace>/private",
			Service:     GetPrivateMessageHandler,
			TokenValidator: jwt.New(
				jwt.WithDebug(),
			).Middleware,
		},
	}
}

func (Y Debug) LongPath(route libgeoip.Route) string {
	return libgeoip.DefaultLongPath(Y, route)
}
