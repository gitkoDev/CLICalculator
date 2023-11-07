package main

import (
	"fmt"
)

const (
	addition       = "add"
	subtraction    = "subtract"
	multiplication = "multiply"
	division       = "divide"
)

func main() {

	for {
		first, second := getOperands()
		operation := getOperation()

		switch operation {
		case addition:
			fmt.Println(add(first, second))
		case subtraction:
			fmt.Println(subtract(first, second))
		case multiplication:
			fmt.Println(multiply(first, second))
		case division:
			fmt.Println(divide(first, second))

		}
	}

}

func getOperands() (int, int) {
	var firstOperand int
	var secondOperand int

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

func add(first, second int) int {
	return first + second
}

func subtract(first, second int) int {
	return first - second
}

func multiply(first, second int) int {
	return first * second
}

func divide(first, second int) int {
	return first / second
}
