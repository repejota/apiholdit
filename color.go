// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"fmt"
	"image/color"
)

// getColor gets a color from a RGB HTML hex string.
func getColor(colorstr string) (color.RGBA, error) {
	var col color.RGBA
	format := "%02x%02x%02x"
	var r, g, b uint8
	_, err := fmt.Sscanf(colorstr, format, &r, &g, &b)
	if err != nil {
		col = color.RGBA{0, 0, 0, 255}
		return col, err
	}
	col = color.RGBA{r, g, b, 255}
	return col, nil
}
