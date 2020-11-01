package triangulate

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type IP net.IP

func (i *IP) UnmarshalJSON(b []byte) error {
	ip := net.ParseIP(strings.ReplaceAll(string(b), "\"", ""))
	if ip == nil {
		return errors.New(fmt.Sprintf("could not parse %s", string(b)))
	}
	tmp := IP(ip)
	*i = tmp
	return nil
}

type Location struct {
	City           string  `json:"city"`
	Longitude      float64 `json:"long"`
	Latitude       float64 `json:"lat"`
	AccuracyRadius uint16  `json:"accuracy_radius"`
}

type PostTriangulateRequest struct {
	IPs []IP `json:"ips"`
}

type PostTriangulateResponse struct {
	IPs   map[string]Location `json:"ips"`
	Error *string             `json:"error,omitempty"`
}
