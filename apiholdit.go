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
}

// NewPlaceHolder ...
func NewPlaceHolder(width int, height int) *PlaceHolder {
	p := PlaceHolder{}
	p.Width = width
	p.Height = height
	p.MarginRatio = 0.2
	rectangle := image.Rect(0, 0, width, height)
	p.Canvas = image.NewRGBA(rectangle)
	return &p
}

// SetBgColor ...
func (p *PlaceHolder) SetBgColor(bgcolor string) error {
	var col color.RGBA
	col, err := getColor(bgcolor)
	if err != nil {
		return err
	}
	p.BackgroundColor = &col
	return nil
}

// SetFgColor ...
func (p *PlaceHolder) SetFgColor(fgcolor string) error {
	var col color.RGBA
	col, err := getColor(fgcolor)
	if err != nil {
		return err
	}
	p.ForegroundColor = &col
	return nil
}

// Render ...
func (p *PlaceHolder) Render() (*bytes.Buffer, error) {
	// Render background
	err := renderBackground(p.Canvas, p.BackgroundColor)
	if err != nil {
		return nil, err
	}

	// Render text
	ttf, err := ioutil.ReadFile("/Users/raul/go/src/github.com/repejota/apiholdit/contrib/Roboto-Black.ttf")
	if err != nil {
		return nil, err
	}
	fontTTF, err := freetype.ParseFont(ttf)
	if err != nil {
		return nil, err
	}
	c := freetype.NewContext()
	c.SetDPI(DPI)
	c.SetFont(fontTTF)
	c.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 0}))
	c.SetDst(p.Canvas)
	rectangle := p.Canvas.Bounds()
	c.SetClip(rectangle)
	c.SetHinting(font.HintingNone)

	text := "Lorem ipsum dolor sit amet."

	// draw with scaled fontsize to get the real text extent
	fontsize, actwidth := maxPointSize(text, c,
		int(float64(p.Width)*(1.0-p.MarginRatio)),
		int(float64(p.Height)*(1.0-p.MarginRatio)))

	actheight := c.PointToFixed(fontsize/2.0) / 64
	xcenter := (float64(p.Width) / 2.0) - (float64(actwidth) / 2.0)
	ycenter := (float64(p.Height) / 2.0) + (float64(actheight) / 2.0)

	// draw the text
	c.SetFontSize(fontsize)
	c.SetSrc(image.NewUniform(p.ForegroundColor))
	_, err = c.DrawString(text, freetype.Pt(int(xcenter), int(ycenter)))
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, p.Canvas)
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

// maxPointSize returns the maximum point size we can use to fit text inside width and height
// as well as the resulting text-width in pixels
func maxPointSize(text string, c *freetype.Context, width, height int) (float64, int) {
	// never let the font size exceed the requested height
	fontsize := 512.00
	for int(c.PointToFixed(fontsize)/64) > height {
		fontsize -= 2
	}

	// find the biggest matching font size for the requested width
	var actwidth int
	for actwidth = width + 1; actwidth > width; fontsize -= 2 {
		c.SetFontSize(fontsize)

		textExtent, err := c.DrawString(text, freetype.Pt(0, 0))
		if err != nil {
			return 0, 0
		}

		actwidth = int(float64(textExtent.X) / 64)
	}

	return fontsize, actwidth
}
