package calculator

import (
	"calculator/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	result := calculator.Add(1, 2)
	if result != 3 {
		t.Error("Expected 3, got ", result)
	}
}
