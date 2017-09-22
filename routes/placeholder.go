// Copyright 2017 The qurl Authors. All rights reserved.

package routes

import (
	"net/http"
	"strconv"

	"github.com/repejota/apiholdit"
)

// PlaceHolder generates an image placeholder.
func PlaceHolder(w http.ResponseWriter, r *http.Request) {
	p := apiholdit.NewPlaceHolder()

	widthstr := r.URL.Query().Get("width")
	width, err := strconv.Atoi(widthstr)
	if err != nil {
		http.Error(w, "Invalid placeholder width", http.StatusBadRequest)
		return
	}
	p.SetWidth(uint(width))

	heightstr := r.URL.Query().Get("height")
	height, err := strconv.Atoi(heightstr)
	if err != nil {
		http.Error(w, "Invalid placeholder height", http.StatusBadRequest)
		return
	}
	p.SetHeight(uint(height))

	bgcolorstr := r.URL.Query().Get("bgcolor")
	err = p.SetBackgroundColor(bgcolorstr)
	if err != nil {
		http.Error(w, "Invalid placeholder background color", http.StatusBadRequest)
		return
	}

	fgcolorstr := r.URL.Query().Get("fgcolor")
	err = p.SetForegroundColor(fgcolorstr)
	if err != nil {
		http.Error(w, "Invalid placeholder foreground color", http.StatusBadRequest)
		return
	}

	err = p.SetText("Lorem ipsum dolor sit amet.")
	if err != nil {
		http.Error(w, "Unable set text", http.StatusInternalServerError)
		return
	}

	err = p.Render()
	if err != nil {
		http.Error(w, "Unable to render image", http.StatusInternalServerError)
		return
	}

	buffer, err := p.EncodePNG()
	if err != nil {
		http.Error(w, "Unable to encode image to PNG format", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	_, err = w.Write(buffer.Bytes())
	if err != nil {
		http.Error(w, "Unable to write image", http.StatusInternalServerError)
		return
	}
}
