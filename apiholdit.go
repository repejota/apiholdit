// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// PlaceHolder ...
type PlaceHolder struct {
	BackgroundColor *color.RGBA
	ForegroundColor *color.RGBA
	Canvas          *image.RGBA
}

// NewPlaceHolder ...
func NewPlaceHolder(width int, height int) *PlaceHolder {
	p := PlaceHolder{}
	rectangle := image.Rect(0, 0, width, height)
	p.Canvas = image.NewRGBA(rectangle)
	return &p
}

// Render ...
func (p *PlaceHolder) Render() {
	draw.Draw(p.Canvas, p.Canvas.Bounds(), &image.Uniform{p.BackgroundColor}, image.ZP, draw.Src)
}

// getColor gets a color from a RGB HTML hex string.
func (p *PlaceHolder) getColor(colorstr string) (color.RGBA, error) {
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

// SetBgColor ...
func (p *PlaceHolder) SetBgColor(bgcolor string) error {
	var col color.RGBA
	col, err := p.getColor(bgcolor)
	if err != nil {
		return err
	}
	p.BackgroundColor = &col
	return nil
}

// SetFgColor ...
func (p *PlaceHolder) SetFgColor(fgcolor string) error {
	var col color.RGBA
	col, err := p.getColor(fgcolor)
	if err != nil {
		return err
	}
	p.ForegroundColor = &col
	return nil
}
