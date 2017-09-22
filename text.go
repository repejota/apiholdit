// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

// getFont ...
func getFont() (*truetype.Font, error) {
	fontPath := "/Users/raul/go/src/github.com/repejota/apiholdit/contrib/Roboto-Black.ttf"

	ttf, err := ioutil.ReadFile(fontPath)
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
func getFontFinalSize(text string, c *freetype.Context, width uint, height uint, marginratio float64) (float64, uint, uint) {
	scaledWidth, scaledHeight := getFontScaledSize(width, height, marginratio)

	finalFontSize := DefaultMaxFontSize

	// find the biggest matching font size for the requested width
	for uint(c.PointToFixed(finalFontSize)/64) > scaledHeight {
		finalFontSize -= 2
	}
	var finalWidth uint
	for finalWidth = width + 1; finalWidth > scaledWidth; finalFontSize -= 2 {
		c.SetFontSize(finalFontSize)
		textExtent, err := c.DrawString(text, freetype.Pt(0, 0))
		if err != nil {
			return 0, 0, 0
		}
		finalWidth = uint(float64(textExtent.X) / 64)
	}
	finalHeight := uint(c.PointToFixed(finalFontSize/2.0) / 64)

	return finalFontSize, finalWidth, finalHeight
}

// getFontScaledSize ...
func getFontScaledSize(width uint, height uint, marginratio float64) (uint, uint) {
	scaledWidth := uint(float64(width) * (1.0 - marginratio))
	scaledHeight := uint(float64(height) * (1.0 - marginratio))
	return scaledWidth, scaledHeight
}
