package operations

import (
	"testing"
)

type TestData struct {
	x, y   float64
	result float64
}

func TestAdd(t *testing.T) {
	testData := []TestData{
		{1, 3, 4},
		{-5, -6, -11},
		{1000, 0, 1000},
		{1, 0.5, 1.5},
	}

	for _, datum := range testData {
		result := datum.x + datum.y

		if result != datum.result {
			t.Errorf("add(%v, %v) FAILED. Want: %v, got %v", datum.x, datum.y, datum.result, result)
		} else {
			t.Logf("add(%v, %v) PASSED. Want: %v, got %v", datum.x, datum.y, datum.result, result)
		}
	}

}

func TestSubtract(t *testing.T) {
	testData := []TestData{
		{1, 3, -2},
		{5, -6, 11},
		{1000, 0, 1000},
		{4, 2.5, 1.5},
	}

	for _, datum := range testData {
		result := datum.x - datum.y

		if result != datum.result {
			t.Errorf("subtract(%v, %v) FAILED. Want: %v, got %v", datum.x, datum.y, datum.result, result)
		} else {
			t.Logf("subtract(%v, %v) PASSED. Want: %v, got %v", datum.x, datum.y, datum.result, result)
		}
	}

}

func TestMultiply(t *testing.T) {
	testData := []TestData{
		{1, 3, 3},
		{5, -6, -30},
		{1000, 1000, 1000000},
		{4, -2.5, -10},
	}

	for _, datum := range testData {
		result := datum.x * datum.y

		if result != datum.result {
			t.Errorf("multiply(%v, %v) FAILED. Want: %v, got %v", datum.x, datum.y, datum.result, result)
		} else {
			t.Logf("multiply(%v, %v) PASSED. Want: %v, got %v", datum.x, datum.y, datum.result, result)
		}
	}
}

func TestDivide(t *testing.T) {
	testData := []TestData{
		{1, 3, 0.3333333333333333},
		{15, -5, -3},
		{1000, 1000, 1},
		{4, 2, 2},
	}

	for _, datum := range testData {
		result := datum.x / datum.y

		if result != datum.result {
			t.Errorf("divide(%v, %v) FAILED. Want: %v, got %v", datum.x, datum.y, datum.result, result)
		} else {
			t.Logf("divide(%v, %v) PASSED. Want: %v, got %v", datum.x, datum.y, datum.result, result)
		}
	}
}
