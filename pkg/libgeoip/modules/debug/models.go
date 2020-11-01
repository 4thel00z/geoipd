package debug

import "geoipd/pkg/libgeoip"

type GetRoutesResponse struct {
	Routes map[string]libgeoip.Route `json:"routes"`
	Error  *string                              `json:"error,omitempty"`
}
