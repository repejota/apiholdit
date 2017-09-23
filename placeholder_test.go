// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import "testing"

func TestNewPlaceHolder(t *testing.T) {
	expectedWidth := DefaultWidth
	expectedHeight := DefaultHeight
	placeholder := NewPlaceHolder()
	if placeholder.Width != expectedWidth {
		t.Fatalf("Placeholder width expected to be %d but got %d", expectedWidth, placeholder.Width)
	}
	if placeholder.Height != expectedHeight {
		t.Fatalf("Placeholder height expected to be %d but got %d", expectedHeight, placeholder.Height)
	}
	if placeholder.Text != DefaultText {
		t.Fatalf("Placehooder text expected to be %s but bot %s", DefaultText, placeholder.Text)
	}
}

func TestSetWidth(t *testing.T) {
	expectedWidth := DefaultWidth
	placeholder := NewPlaceHolder()
	if placeholder.Width != expectedWidth {
		t.Fatalf("Placeholder width expected to be %d but got %d", expectedWidth, placeholder.Width)
	}
	newWidth := int(1024)
	err := placeholder.SetWidth(newWidth)
	if err != nil {
		t.Fatal(err)
	}
	if placeholder.Width != newWidth {
		t.Fatalf("Placeholder width expected to be %d but got %d", newWidth, placeholder.Width)
	}
}

func TestSetWidthNegativeNumber(t *testing.T) {
	expectedWidth := DefaultWidth
	placeholder := NewPlaceHolder()
	if placeholder.Width != expectedWidth {
		t.Fatalf("Placeholder width expected to be %d but got %d", expectedWidth, placeholder.Width)
	}
	newWidth := int(-640)
	err := placeholder.SetWidth(newWidth)
	if err.Error() != "width must be >= 0" {
		t.Fatalf("Expected to fail because it is a negative number but but got %s", err)
	}
}

func TestSetHeight(t *testing.T) {
	expectedHeight := DefaultHeight
	placeholder := NewPlaceHolder()
	if placeholder.Height != expectedHeight {
		t.Fatalf("Placeholder height expected to be %d but got %d", expectedHeight, placeholder.Width)
	}
	newHeight := int(768)
	placeholder.SetHeight(newHeight)
	if placeholder.Height != newHeight {
		t.Fatalf("Placeholder height expected to be %d but got %d", newHeight, placeholder.Width)
	}
}

func TestSetHeightNegativeNumber(t *testing.T) {
	expectedHeight := DefaultHeight
	placeholder := NewPlaceHolder()
	if placeholder.Height != expectedHeight {
		t.Fatalf("Placeholder height expected to be %d but got %d", expectedHeight, placeholder.Height)
	}
	newHeight := int(-768)
	err := placeholder.SetHeight(newHeight)
	if err.Error() != "height must be >= 0" {
		t.Fatalf("Expected to fail because it is a negative number but but got %s", err)
	}
}

func TestSetText(t *testing.T) {
	placeholder := NewPlaceHolder()
	if placeholder.Text != DefaultText {
		t.Fatalf("Placehooder text expected to be %s but bot %s", DefaultText, placeholder.Text)
	}
	expectedText := "foo bar"
	placeholder.SetText(expectedText)
	if placeholder.Text != expectedText {
		t.Fatalf("Placehooder text expected to be %s but bot %s", expectedText, placeholder.Text)
	}
}

func TestSetBgColor(t *testing.T) {
	placeholder := NewPlaceHolder()
	err := placeholder.SetBackgroundColor("ff0000")
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if placeholder.BackgroundColor.R != 255 {
		t.Fatalf("Placeholder background color R expected to be 255 but got %d", placeholder.BackgroundColor.R)
	}
	if placeholder.BackgroundColor.G != 0 {
		t.Fatalf("Placeholder background color R expected to be 0 but got %d", placeholder.BackgroundColor.G)
	}
	if placeholder.BackgroundColor.B != 0 {
		t.Fatalf("Placeholder background color R expected to be 0 but got %d", placeholder.BackgroundColor.B)
	}
	if placeholder.BackgroundColor.A != 255 {
		t.Fatalf("Placeholder background color A expected to be 255 but got %d", placeholder.BackgroundColor.A)
	}
}

func TestSetBgColorFail(t *testing.T) {
	placeholder := NewPlaceHolder()
	err := placeholder.SetBackgroundColor("foo")
	if err == nil {
		t.Fatalf("Expected to error but got %s", err)
	}
}

func TestSetFgColor(t *testing.T) {
	placeholder := NewPlaceHolder()
	err := placeholder.SetForegroundColor("00ffff")
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if placeholder.ForegroundColor.R != 0 {
		t.Fatalf("Placeholder foreground color R expected to be 0 but got %d", placeholder.ForegroundColor.R)
	}
	if placeholder.ForegroundColor.G != 255 {
		t.Fatalf("Placeholder foreground color R expected to be 255 but got %d", placeholder.ForegroundColor.G)
	}
	if placeholder.ForegroundColor.B != 255 {
		t.Fatalf("Placeholder foreground color R expected to be 255 but got %d", placeholder.ForegroundColor.B)
	}
	if placeholder.ForegroundColor.A != 255 {
		t.Fatalf("Placeholder foreground color A expected to be 255 but got %d", placeholder.ForegroundColor.A)
	}
}

func TestSetFgColorFail(t *testing.T) {
	placeholder := NewPlaceHolder()
	err := placeholder.SetForegroundColor("foo")
	if err == nil {
		t.Fatalf("Expected to error but got %s", err)
	}
}
