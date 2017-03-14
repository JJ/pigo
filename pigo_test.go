package pigo
// Taken from https://play.golang.org/p/mb5eoZpnYN
// No license specified

import (
	"testing"
)


func TestSmall(t *testing.T) {
	t.Log("Test for smaller than 7")
	if digits := Pi(3); digits != "3.141595" {
		t.Error("Expected 3.141595, but it was %s instead", digits)
	}
}

func TestBig(t *testing.T) {
	t.Log("Test for a bit bigger")
	if digits := Pi(10); digits != "3.141592653" {
		t.Error("Expected 3.141592653, but it was %s instead", digits)
	}
}
