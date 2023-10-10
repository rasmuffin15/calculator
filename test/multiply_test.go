package calculator

import (
	"calculator/calculator"
	"testing"
)

func TestMultiply(t *testing.T) {
	result := calculator.Multiply(1, 2)
	if result != 2 {
		t.Error("Expected 2, got ", result)
	}
}
