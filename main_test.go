package main

import "testing"

func TestReturnSome(t *testing.T) {
	if ReturnSome() != "apekatt" {
		t.Errorf("error: was not gris")
	}
}
