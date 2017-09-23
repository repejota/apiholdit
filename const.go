// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import "image/color"

const (
	// DefaultServerAddress is the default hostname where the server is going
	// to be started and listening.
	DefaultServerAddress = "localhost"

	// DefaultServerPort is the default port number where the server is going
	// to be started and listening.
	DefaultServerPort = "8080"

	// DefaultWidth is the default width of the image placeholder.
	DefaultWidth = 640

	// DefaultHeight is the default height of the image placeholder.
	DefaultHeight = 480

	// DefaultDPI is the default DPI value for the image placeholder.
	DefaultDPI = 72.00

	// DefaultMarginRatio is the default margin ratio to be used to render
	// the text on the image placeholder.
	//
	// This percentage margin will be applied to both four inner margin of the
	// image bounds.
	DefaultMarginRatio = 0.2

	// DefaultMaxFontSize is the maximum size of the font that will be used to
	// render the placeholder text.
	DefaultMaxFontSize = 512.00
)

var (
	// DefaultBackgroundColor : Silver #bdc3c7
	// http://flatuicolors.com/
	DefaultBackgroundColor = &color.RGBA{189, 195, 199, 255}

	// DefaultForegroundColor : Midnight Blue #2c3e50
	// http://flatuicolors.com/
	DefaultForegroundColor = &color.RGBA{44, 62, 80, 255}
)
