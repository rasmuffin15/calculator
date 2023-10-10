package calculator

import (
	"calculator/calculator"
	"testing"
)

func TestDivide(t *testing.T) {
	result := calculator.Divide(1, 2)
	if result != 0 {
		t.Error("Expected 0, got ", result)
	}
}
