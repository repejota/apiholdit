// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"

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
func NewPlaceHolder() *PlaceHolder {
	p := PlaceHolder{}
	p.Width = DefaultWidth
	p.Height = DefaultHeight
	p.MarginRatio = DefaultMarginRatio
	rectangle := image.Rect(0, 0, DefaultWidth, DefaultHeight)
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

// EncodeGIF ...
func (p *PlaceHolder) EncodeGIF(options *gif.Options) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	err := gif.Encode(buffer, p.Canvas, options)
	if err != nil {
		return buffer, nil
	}
	return buffer, err
}

// EncodeJPEG ...
func (p *PlaceHolder) EncodeJPEG(options *jpeg.Options) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	err := jpeg.Encode(buffer, p.Canvas, options)
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
