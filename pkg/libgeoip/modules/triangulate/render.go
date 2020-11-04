package triangulate

import (
	"github.com/flopp/go-staticmaps"
	"github.com/golang/geo/s2"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
)

func RenderPng(writer io.Writer, width, height int, lat, lng, size float64, markerColor color.RGBA) error {
	ctx := sm.NewContext()
	ctx.SetSize(width, height)
	ctx.AddMarker(sm.NewMarker(s2.LatLngFromDegrees(lat, lng), markerColor, size))
	img, err := ctx.Render()
	if err != nil {
		return err
	}

	return png.Encode(writer, img)
}

func RenderJpeg(writer io.Writer, width, height int, lat, lng, size float64, markerColor color.RGBA, options *jpeg.Options) error {
	ctx := sm.NewContext()
	ctx.SetSize(width, height)
	ctx.AddMarker(sm.NewMarker(s2.LatLngFromDegrees(lat, lng), markerColor, size))
	img, err := ctx.Render()
	if err != nil {
		return err
	}

	return jpeg.Encode(writer, img, options)
}
