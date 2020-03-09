package main

import "testing"

func TestReturnSome(t *testing.T) {
	if ReturnSome() != "gris" {
		t.Errorf("error: was not gris")
	}
}
