package debug

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/monzo/typhon"
	"geoipd/pkg/libgeoip"
	libjwt "geoipd/pkg/libgeoip/jwt"
)

func GetRoutesHandler(app libgeoip.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {

		response := req.Response(&GetRoutesResponse{
			Routes: app.Routes(),
		})

		response.StatusCode = 200
		return response
	}
}

func GetPrivateMessageHandler(app libgeoip.App) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		token := req.Value(libjwt.DefaultUserProperty).(*jwt.Token)
		response := req.Response(&libgeoip.GenericResponse{
			Message: fmt.Sprintf("This is my token: %s!", token.Raw),
		})
		response.StatusCode = 200
		return response
	}
}
