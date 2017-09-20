// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/repejota/apiholdit"
)

// renderImage renders an image of specified size.
func renderImage(width int, height int, bgcolor *color.RGBA, fgcolor *color.RGBA) image.Image {
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{bgcolor}, image.ZP, draw.Src)
	var img image.Image = canvas
	return img
}

// writeImage encodes an image 'img' in PNG format and writes it
// into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

// PlaceHolder generates an image placeholder.
func PlaceHolder(w http.ResponseWriter, r *http.Request) {
	p := apiholdit.PlaceHolder{}

	width, err := strconv.Atoi(r.URL.Query().Get("width"))
	if err != nil {
		http.Error(w, "Invalid image width", http.StatusBadRequest)
		return
	}
	height, err := strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil {
		http.Error(w, "Invalid image height", http.StatusBadRequest)
		return
	}

	bgcolor, err := p.SetBgColor(r.URL.Query().Get("bgcolor"))
	if err != nil {
		http.Error(w, "Invalid background color", http.StatusBadRequest)
		return
	}
	fgcolor, err := p.SetFgColor(r.URL.Query().Get("fgcolor"))
	if err != nil {
		http.Error(w, "Invalid foreground color", http.StatusBadRequest)
		return
	}

	img := renderImage(width, height, &bgcolor, &fgcolor)
	writeImage(w, &img)
}
