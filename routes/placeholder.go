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
)

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

// renderImage renders an image of specified size.
func renderImage(width int, height int) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	return img
}

// PlaceHolder generates an image placeholder.
func PlaceHolder(w http.ResponseWriter, r *http.Request) {
	width := 640
	height := 480

	img := renderImage(width, height)
	writeImage(w, &img)
}
