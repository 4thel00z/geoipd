package filters

import (
	"github.com/monzo/typhon"
	"geoipd/pkg/libgeoip"
)

func Auth(app libgeoip.App) typhon.Filter {
	return func(req typhon.Request, svc typhon.Service) typhon.Response {
		pattern := app.Router.Pattern(req)
		routes := app.Routes()
		route, ok := routes[pattern]
		if !ok {
			return svc(req)
		}

		if route.TokenValidator == nil {
			return svc(req)
		}

		validator := route.TokenValidator

		if validator == nil {
			return svc(req)
		}

		return validator(req, svc)

	}
}
