package main

import (
	"testing"
)

func TestDevide(t *testing.T) {
	_, err := Devide(10.0, 1.0)

	if err != nil {
		t.Error("Yah error padahal harusnya sih nggak error")
	}
}
