package operations

import "testing"

var (
	firstOperand  float64 = 11
	secondOperand float64 = 5.5
)

func TestAdd(t *testing.T) {
	result := firstOperand + secondOperand

	if result != 16.5 {
		t.Errorf("Add(11, 5.5) FAIL. Want: %v, got: %v", 16.5, result)
	}
}

func TestSubtract(t *testing.T) {
	result := firstOperand - secondOperand

	if result != 5.5 {
		t.Errorf("Subtract(11, 5.5) FAIL. Want: %v, got: %v", 5.5, result)
	}
}

func TestMultiply(t *testing.T) {
	result := firstOperand * secondOperand

	if result != 60.5 {
		t.Errorf("Mutiply(11, 5.5) FAIL. Want: %v, got: %v", 60.5, result)
	}
}

func TestDivide(t *testing.T) {
	result := firstOperand / secondOperand

	if result != 2 {
		t.Errorf("Divide(11, 5.5) FAIL. Want: %v, got: %v", 2, result)
	}
}
