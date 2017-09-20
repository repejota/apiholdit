// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/repejota/apiholdit"
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

	p := apiholdit.NewPlaceHolder(width, height)

	err = p.SetBgColor(r.URL.Query().Get("bgcolor"))
	if err != nil {
		http.Error(w, "Invalid background color", http.StatusBadRequest)
		return
	}
	err = p.SetFgColor(r.URL.Query().Get("fgcolor"))
	if err != nil {
		http.Error(w, "Invalid foreground color", http.StatusBadRequest)
		return
	}

	p.Render()
	var img image.Image
	img = p.Canvas
	writeImage(w, &img)
}
