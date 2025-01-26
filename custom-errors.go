package main

import (
	"errors"
	"fmt"
)

// custom errors usually have the ending "Error"
type argError struct {
	arg     int
	message string
}

// adding this error method makes argError implement the error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// returning custom error itself
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
