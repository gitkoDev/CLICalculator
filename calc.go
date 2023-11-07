package main

import (
	"CLICalculator/operations"
	"fmt"
)

const (
	addition       = "add"
	subtraction    = "subtract"
	multiplication = "multiply"
	division       = "divide"
)

func main() {

	first, second := getOperands()
	operation := getOperation()

	switch operation {
	case addition:
		operations.Add(first, second)
	case subtraction:
		operations.Subtract(first, second)
	case multiplication:
		operations.Multiply(first, second)
	case division:
		operations.Divide(first, second)
	}

}

func getOperands() (float64, float64) {
	var firstOperand float64
	var secondOperand float64

	fmt.Println("Number 1:")
	fmt.Scan(&firstOperand)
	fmt.Println("Number 2:")
	fmt.Scan(&secondOperand)

	return firstOperand, secondOperand
}

func getOperation() string {
	var operation string

	fmt.Println("What operation would you like to perform?")
	fmt.Scan(&operation)

	switch operation {
	case addition, subtraction, multiplication, division:
		return operation
	default:
		return "Unknown operation"
	}

}
