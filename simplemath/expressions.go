package simplemath

import (
	"errors"
)

// func main() {
// 	result, err := simplemath.Divide(6, 0)
// 	if err != nil {
// 		fmt.Printf("An error occured %s\n", err.Error())
// 	} else {
// 		fmt.Printf("%f \n", result)
// 	}

// 	numbers := []float64{12.2, 14, 16, 22.4}
// 	total := simplemath.Sum(numbers...) // taking a slice and spreading them out.
// 	fmt.Printf("total of sum: %f\n", total)
// }

// Sum ...
func Sum(values ...float64) float64 { // taking a bunch of numbers and putting them in slice.
	total := 0.0
	for _, value := range values {
		total = total + value
	}

	return total
}

// Divide ...
func Divide(p1, p2 float64) (answer float64, err error) {
	if p2 == 0 {
		err = errors.New("cannot divide by zero")
	}

	answer = p1 / p2
	return
}

// Add ...
func Add(p1, p2 float64) float64 {
	return p1 + p2
}

// Subtract ...
func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

// Multiply ...
func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}
