// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// PlaceHolder ...
type PlaceHolder struct {
	Width           int
	Height          int
	MarginRatio     float64
	BackgroundColor *color.RGBA
	ForegroundColor *color.RGBA
	Canvas          *image.RGBA
	Text            string
}

// NewPlaceHolder ...
func NewPlaceHolder(width int, height int) *PlaceHolder {
	p := PlaceHolder{}
	p.Width = width
	p.Height = height
	p.MarginRatio = DefaultMarginRatio
	rectangle := image.Rect(0, 0, width, height)
	p.Canvas = image.NewRGBA(rectangle)
	return &p
}

// SetBackgroundColor ...
func (p *PlaceHolder) SetBackgroundColor(bgcolor string) error {
	var col color.RGBA
	col, err := getColor(bgcolor)
	if err != nil {
		return err
	}
	p.BackgroundColor = &col
	return nil
}

// SetForegroundColor ...
func (p *PlaceHolder) SetForegroundColor(fgcolor string) error {
	var col color.RGBA
	col, err := getColor(fgcolor)
	if err != nil {
		return err
	}
	p.ForegroundColor = &col
	return nil
}

// SetText ...
func (p *PlaceHolder) SetText(text string) error {
	p.Text = text
	return nil
}

// Render ...
func (p *PlaceHolder) Render() error {
	// Render background
	err := renderBackground(p.Canvas, p.BackgroundColor)
	if err != nil {
		return err
	}

	// Get font to be used
	fontTTF, err := getFont()
	if err != nil {
		return err
	}

	// Render text
	err = renderText(p.Canvas, fontTTF, p.Width, p.Height, p.MarginRatio, p.Text, p.ForegroundColor)
	if err != nil {
		return err
	}

	return nil
}

// EncodePNG ...
func (p *PlaceHolder) EncodePNG() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, p.Canvas)
	if err != nil {
		return buffer, nil
	}
	return buffer, err
}

// renderBackground ...
func renderBackground(canvas *image.RGBA, bgcolor *color.RGBA) error {
	rectangle := canvas.Bounds()
	color := &image.Uniform{bgcolor}
	draw.Draw(canvas, rectangle, color, image.ZP, draw.Src)
	return nil
}

// renderText ...
func renderText(canvas *image.RGBA, fontTTF *truetype.Font, width int, height int, marginratio float64, text string, fgcolor *color.RGBA) error {
	rectangle := canvas.Bounds()

	context := freetype.NewContext()
	context.SetDPI(DefaultDPI)
	context.SetFont(fontTTF)
	context.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 0}))
	context.SetDst(canvas)
	context.SetClip(rectangle)
	context.SetHinting(font.HintingNone)

	// calculate final font size and coordinates to draw
	finalFontSize, finalWidth, finalHeight := getFontFinalSize(text, context, width, height, marginratio)
	context.SetFontSize(finalFontSize)

	// draw the text
	context.SetSrc(image.NewUniform(fgcolor))
	xCenter := (float64(width) / 2.0) - (float64(finalWidth) / 2.0)
	yCenter := (float64(height) / 2.0) + (float64(finalHeight) / 2.0)
	_, err := context.DrawString(text, freetype.Pt(int(xCenter), int(yCenter)))
	if err != nil {
		return err
	}

	return nil
}

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

// getColor gets a color from a RGB HTML hex string.
func getColor(colorstr string) (color.RGBA, error) {
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
