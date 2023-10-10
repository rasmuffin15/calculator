package calculator

import (
	"calculator/calculator"
	"testing"
)

func TestSubtract(t *testing.T) {
	result := calculator.Subtract(1, 2)
	if result != -1 {
		t.Error("Expected -1, got ", result)
	}
}
