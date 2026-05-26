package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//io package provides basic interfaces and functions of I/O operations
//io has two interfaces Reader and Writer
//type Reader interface{ Read(p []byte) error} and type Writer interface{ Write(p []byte) (n int,e err)}

//any type which implements these methods Read and Write can be used anywhere a Reader or writer is expected like
// a file is reader or writer
// http response body

func inputOutput() {
	fmt.Println("io.go")

	//creates a new Reader
	r := strings.NewReader("hello")

	//this would load everything in memory so not for large files
	data, err := io.ReadAll(r)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(data))

	r1 := strings.NewReader("don't know\n")

	n, err := io.Copy(os.Stdout, r1)

	if err != nil {
		fmt.Println("error is", err)
		return
	}
	fmt.Println("bytes copied :", n)

	//io.WriteString
	n1, err1 := io.WriteString(os.Stdout, "sorry! and thank you")

	if err1 != nil {
		fmt.Println("error occured", err1)
		return
	}
	fmt.Println(n1)

	//io.LimitReader
	r3 := strings.NewReader("don't know\n")
	limitedReader := io.LimitReader(r3, 5)

	n4, err6 := io.Copy(os.Stdout, limitedReader)
	if err6 != nil {
		fmt.Println("error :", err6)
		return
	}
	fmt.Println(n4)

}
