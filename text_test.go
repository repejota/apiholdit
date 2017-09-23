// Copyright 2017 The apiholdit Authors. All rights reserved.

package apiholdit

import (
	"testing"
)

func TestGetFont(t *testing.T) {
	_, err := getFont()
	if err != nil {
		t.Fatal(err)
	}
}
