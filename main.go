package main

import (
	"errors"
	"fmt"
	"io"
	"mrigankapaul/simplemath"
)

type mathExpr = string

const (
	addExpr      = mathExpr("add")
	subtractExpr = mathExpr("subtract")
	multiplyExpr = mathExpr("multiply")
)

func main() {
	if err := ReadFullFile(); err != nil {
		fmt.Printf("something bad occurred: %s\n", err)
	}
}

func mathExpression(expr mathExpr) func(float64, float64) float64 {
	switch expr {
	case addExpr:
		return simplemath.Add
	case subtractExpr:
		return simplemath.Subtract
	case multiplyExpr:
		return simplemath.Multiply
	default:
		panic("an invalid math expression was used")
	}
}

// ReadFullFile ...
func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}

	defer func() {
		_ = r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("A panic occurred but is OK")
		} else if p != nil {
			panic("an unexpected  error happened and we dont want to recover")
		}
	}()

	defer func() {
		println("Before for loop")
	}()

	for {
		value, readerErr := r.Read([]byte("text that does something"))

		if err == io.EOF {
			println("finished reading the file")

			break
		} else if err != nil {
			err = readerErr
			return
		}

		println(value)
	}

	defer func() {
		println("After for loop")
	}()

	return nil
}

// SimpleReader ...
type SimpleReader struct {
	count int
}

var errCatastrophicReader = errors.New("somethig catastophic happened")

func (br *SimpleReader) Read(p []byte) (n int, err error) {

	if br.count == 2 {
		//  panic(errCatastrophicReader)
		panic("Another Catastropic error")
	}
	if br.count > 3 {
		return 0, io.EOF
		// return 0, errors.New("random error")
	}

	br.count++
	return br.count, nil
}

// Close ...
func (br *SimpleReader) Close() error {
	println("Closing Reader")
	return nil
}
