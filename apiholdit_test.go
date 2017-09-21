// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import "testing"

func TestNewPlaceHolder(t *testing.T) {
	placeholder := NewPlaceHolder(640, 480)
	if placeholder.Width != 640 {
		t.Fatalf("Placeholder width expected to be 640 but got %d", placeholder.Width)
	}
	if placeholder.Height != 480 {
		t.Fatalf("Placeholder height expected to be 480 but got %d", placeholder.Height)
	}
}

func TestSetBgColort(t *testing.T) {
	placeholder := NewPlaceHolder(640, 480)
	placeholder.SetBgColor("ff0000")
	if placeholder.BackgroundColor.R != 255 {
		t.Fatalf("Placeholder background color R expected to be 255 but got %d", placeholder.BackgroundColor.R)
	}
}
