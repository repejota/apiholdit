// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"testing"
)

func TestGetFont(t *testing.T) {
	_, err := getFont(DefaultFontPath)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFontNotFound(t *testing.T) {
	path := "/tmp/invalidfont.ttf"
	_, err := getFont(path)
	if err.Error() != "open /tmp/invalidfont.ttf: no such file or directory" {
		t.Fatal(err)
	}
}
