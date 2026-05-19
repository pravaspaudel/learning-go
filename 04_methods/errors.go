package main

import "fmt"

//go express error state with error values
//in go, error type is built-in interface

//go internally has this interface
type error interface {
	Error() string
}

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, MyError{"cannot divided by zero"}
	}
	return a / b, nil
}

func errors() {
	fmt.Println("this is errors.go")

	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Result : ", result)
}
