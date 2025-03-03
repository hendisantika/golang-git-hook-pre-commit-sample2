package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Multiply(2, 3) = %d; want %d", result, expected)
	}
}
