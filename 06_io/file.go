package main

import (
	"fmt"
	"io"
	"os"
)

func files() {
	fmt.Println("files.go")

	// file, err := os.Create("first.txt")

	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return
	// }
	// fmt.Println("successfully created file")

	// content := "add this content to the file"

	// fmt.Println("writting to the file.....")

	// io.WriteString(file, content)

	// defer file.Close()

	//reading from file
	file, err := os.Open("first.txt")
	if err != nil {
		fmt.Println("error while reading the file : ", file)
	}

	defer file.Close()

	buff := make([]byte, 1024)

	for {
		_, err := file.Read(buff)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("error while reading the file", err)
			return
		}
		fmt.Println(string(buff))
	}

}
