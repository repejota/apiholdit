// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"fmt"
	"image/color"
)

// PlaceHolder ...
type PlaceHolder struct {
	BackgroundColor *color.RGBA
	ForegroundColor *color.RGBA
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
func (p *PlaceHolder) SetBgColor(bgcolor string) (color.RGBA, error) {
	var col color.RGBA
	col, err := p.getColor(bgcolor)
	if err != nil {
		col, err = p.getColor("000000")
		return col, err
	}
	return col, nil
}

// SetFgColor ...
func (p *PlaceHolder) SetFgColor(fgcolor string) (color.RGBA, error) {
	var col color.RGBA
	col, err := p.getColor(fgcolor)
	if err != nil {
		col, err = p.getColor("ffffff")
		return col, err
	}
	return col, nil
}
