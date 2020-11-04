package triangulate

import (
	"bytes"
	"errors"
	"fmt"
	"geoipd/pkg/libgeoip"
	"geoipd/pkg/libgeoip/filters"
	"github.com/monzo/typhon"
	"net"
	"net/http"
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
		location, ok := Locations(app, tReq.IP)[net.IP.String(net.IP(tReq.IP))]
		if !ok {
			response := req.Response(errors.New(fmt.Sprintf("couldn't find geo entry for %s ", tReq.IP)))
			response.StatusCode = http.StatusNotFound
			return response
		}
		buffer := bytes.NewBuffer([]byte{})
		color, err := ParseHexColor(tReq.Color)
		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusBadRequest
			return response
		}

		err = RenderPng(buffer, tReq.Width, tReq.Height, location.Latitude, location.Longitude, tReq.Size, color)
		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}

		response := req.Response(nil)
		response.Header.Add("Content-Type", "image/png")
		_, err = response.Write(buffer.Bytes())
		if err != nil {
			response := req.Response(err.Error())
			response.StatusCode = http.StatusInternalServerError
			return response
		}
		response.StatusCode = 200
		return response
	}
}
