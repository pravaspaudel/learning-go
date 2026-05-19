package main

import (
	"fmt"
	"io"
	"strings"
)

type Number interface {
	int | float64
}

func add[T Number](a, b T) T {
	return a + b
}

func inputs() {
	fmt.Println("\ninput.go")

	x, y := 4.5, 3.0
	fmt.Println(add(x, y))

	// var Name string

	// fmt.Print("enter your name : ")
	// fmt.Scan(&Name)

	// fmt.Println("hello", Name)

	// var a, b int
	// fmt.Println("enter value of a and b : ")
	// fmt.Scan(&a, &b)
	// fmt.Printf("sum of %d , %d is %d", a, b, a+b)

	//buffered input
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("enter text :")
	// text, _ := reader.ReadString('\n')
	// fmt.Println("you typed :", text)

	r := strings.NewReader("hello, Reader !")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
