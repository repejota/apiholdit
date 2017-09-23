// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"fmt"
	"image/color"
)

// getColor gets a color from a RGB HTML hex string.
func getColor(colorstr string) (*color.RGBA, error) {
	var r, g, b uint8
	format := "%02x%02x%02x"
	_, err := fmt.Sscanf(colorstr, format, &r, &g, &b)
	if err != nil {
		return DefaultBackgroundColor, err
	}
	return &color.RGBA{r, g, b, 255}, nil
}
