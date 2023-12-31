package operations

import (
	"errors"
	"fmt"
)

func Add(first, second float64) {
	fmt.Println(first + second)
}

func Subtract(first, second float64) {
	fmt.Println(first - second)
}
func Multiply(first, second float64) {
	fmt.Println(first * second)
}

func Divide(first, second float64) error {
	if second == 0 {
		return errors.New("cannot divide by zero")
	}
	fmt.Println(first / second)
	return nil
}
