// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

// getFont ...
func getFont(path string) (*truetype.Font, error) {
	ttf, err := Asset(path)
	if err != nil {
		return nil, err
	}

	fontTTF, err := freetype.ParseFont(ttf)
	if err != nil {
		return nil, err
	}

	return fontTTF, nil
}

// maxPointSize returns the maximum point size we can use to fit text inside
// width and height as well as the resulting text-width in pixels
func getFontFinalSize(text string, c *freetype.Context, width int, height int, marginratio float64) (float64, int, int) {
	scaledWidth, scaledHeight := getFontScaledSize(width, height, marginratio)

	finalFontSize := DefaultMaxFontSize

	// find the biggest matching font size for the requested width
	for int(c.PointToFixed(finalFontSize)/64) > scaledHeight {
		finalFontSize -= 2
	}
	var finalWidth int
	for finalWidth = width + 1; finalWidth > scaledWidth; finalFontSize -= 2 {
		c.SetFontSize(finalFontSize)
		textExtent, err := c.DrawString(text, freetype.Pt(0, 0))
		if err != nil {
			return 0, 0, 0
		}
		finalWidth = int(float64(textExtent.X) / 64)
	}
	finalHeight := int(c.PointToFixed(finalFontSize/2.0) / 64)

	return finalFontSize, finalWidth, finalHeight
}

// getFontScaledSize ...
func getFontScaledSize(width int, height int, marginratio float64) (int, int) {
	scaledWidth := int(float64(width) * (1.0 - marginratio))
	scaledHeight := int(float64(height) * (1.0 - marginratio))
	return scaledWidth, scaledHeight
}
