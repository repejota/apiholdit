// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"bytes"
	"fmt"
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
func renderImage(width int, height int, color *color.RGBA) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)

	var img image.Image = m
	return img
}

// getColor gets a color from a RGB HTML hex string.
func getColor(colorstr string) (color.RGBA, error) {
	fmt.Println(colorstr)

	var col color.RGBA
	format := "%02x%02x%02x"
	var r, g, b uint8
	n, err := fmt.Sscanf(colorstr, format, &r, &g, &b)
	if err != nil {
		col = color.RGBA{0, 0, 0, 255}
		return col, err
	}
	if n != 3 {
		col = color.RGBA{0, 0, 0, 255}
		return col, fmt.Errorf("color: %v is not a hex-color", colorstr)
	}
	col = color.RGBA{r, g, b, 255}
	return col, nil
}

// PlaceHolder generates an image placeholder.
func PlaceHolder(w http.ResponseWriter, r *http.Request) {

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

	bgcolor, err := getColor(r.URL.Query().Get("bgcolor"))
	if err != nil {
		http.Error(w, "Invalid color", http.StatusBadRequest)
		return
	}

	img := renderImage(width, height, &bgcolor)
	writeImage(w, &img)
}
