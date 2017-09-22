// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import "testing"

func TestNewPlaceHolder(t *testing.T) {
	expectedWidth := uint(640)
	expectedHeight := uint(480)
	placeholder := NewPlaceHolder()
	if placeholder.Width != expectedWidth {
		t.Fatalf("Placeholder width expected to be %d but got %d", expectedWidth, placeholder.Width)
	}
	if placeholder.Height != expectedHeight {
		t.Fatalf("Placeholder height expected to be %d but got %d", expectedHeight, placeholder.Height)
	}
}

func TestSetWidth(t *testing.T) {
	expectedWidth := uint(640)
	placeholder := NewPlaceHolder()
	if placeholder.Width != expectedWidth {
		t.Fatalf("Placeholder width expected to be %d but got %d", expectedWidth, placeholder.Width)
	}
	newWidth := uint(1024)
	placeholder.SetWidth(newWidth)
	if placeholder.Width != newWidth {
		t.Fatalf("Placeholder width expected to be %d but got %d", newWidth, placeholder.Width)
	}
}

func TestSetHeight(t *testing.T) {
	expectedHeight := uint(480)
	placeholder := NewPlaceHolder()
	if placeholder.Height != expectedHeight {
		t.Fatalf("Placeholder height expected to be %d but got %d", expectedHeight, placeholder.Width)
	}
	newHeight := uint(768)
	placeholder.SetHeight(newHeight)
	if placeholder.Height != newHeight {
		t.Fatalf("Placeholder height expected to be %d but got %d", newHeight, placeholder.Width)
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
