// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"testing"
)

func TestGetFont(t *testing.T) {
	_, err := getFont("data/fonts/Roboto-Black.ttf")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFontNotFound(t *testing.T) {
	path := "/tmp/invalidfont.ttf"
	_, err := getFont(path)
	if err.Error() != "Asset /tmp/invalidfont.ttf not found" {
		t.Fatal(err)
	}
}

func TestGetFontInvalid(t *testing.T) {
	path := "data/testdata/invalidfont.ttf"
	_, err := getFont(path)
	if err.Error() != "freetype: invalid TrueType format: bad TTF version" {
		t.Fatal(err)
	}
}

func TestGetFontScaledSize(t *testing.T) {
	expectedScaledWidth := 512
	expectedScaledHeight := 384
	scaledWidth, scaledHeight := getFontScaledSize(DefaultWidth, DefaultHeight, DefaultMarginRatio)
	if scaledWidth != expectedScaledWidth {
		t.Fatalf("Text scaledWith expected to be %d but got %d", expectedScaledWidth, scaledWidth)
	}
	if scaledHeight != expectedScaledHeight {
		t.Fatalf("Text scaledHeight expected to be %d but got %d", expectedScaledHeight, scaledHeight)
	}
}
