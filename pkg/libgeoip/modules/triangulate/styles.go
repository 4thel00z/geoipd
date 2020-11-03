package triangulate

import (
	"errors"
	"fmt"
	"github.com/paulmach/orb/geojson"
)

type StyleMap map[string]map[string]interface{}

var (
	Styles = StyleMap{
		"Polygon": {
			"fill":         "red",
			"stroke-width": "3",
			"fill-opacity": 0.6,
		},
		"MultiPolygon": {
			"fill":         "red",
			"stroke-width": "3",
			"fill-opacity": 0.6,
		},
		"LineString": {
			"fill":         "red",
			"stroke-width": "3",
			"fill-opacity": 0.6,
		},
		"MultiLineString": {
			"fill":         "red",
			"stroke-width": "3",
			"fill-opacity": 0.6,
		},
		"Point": {
			"fill":         "blue",
			"stroke-width": "3",
			"fill-opacity": 0.6,
		},
		"MultiPoint": {
			"fill":         "red",
			"stroke-width": "3",
			"fill-opacity": 0.6,
		},
		"GeometryCollection": {
			"fill":         "red",
			"stroke-width": "3",
			"fill-opacity": 0.6,
		},
	}
)

func (s StyleMap) Keys() []string {
	keys := []string{}
	for _, v := range s {
		for k, _ := range v {
			keys = append(keys, k)
		}
	}

	return keys
}

func (s StyleMap) AddToFeature(f *geojson.Feature) error {
	t := f.Geometry.GeoJSONType()
	style, ok := s[t]
	if !ok {
		return errors.New(fmt.Sprintf("GeoJSONType %s not supported", t))
	}
	for k, v := range style {
		f.Properties[k] = v
	}
	return nil
}
