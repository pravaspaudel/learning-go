package main

import "fmt"

type vertex struct {
	X int
	Y int
}

func structs() {
	v := vertex{
		X: 32,
		Y: 22,
	}
	var v2 vertex
	v2 = vertex{92, 89}

	v3 := vertex{}      //both have zero values
	v4 := vertex{X: 82} //both have zero values

	fmt.Println(v)
	fmt.Println(v2.X, v2.Y)
	fmt.Println(v3.X, v3.Y)
	fmt.Println(v4.X, v4.Y)
}

func arrays() {
	//array defined as
	ar := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//slices - slices are reference to the arrays
	//zero value of slice [] is nil
	var sl []int
	sl = ar[4:8]

	ar[5] = 50

	for i, val := range sl {
		fmt.Println(i, val)
	}

	//slices of structures
	slstruct := []Rect{{1, 2}, {4, 5}, {6, 7}, {8, 9}}

	//length and capacity of slice
	slstruct = append(slstruct, Rect{99, 99})
	fmt.Println(slstruct)
	fmt.Printf("length of slstruct is %d and capacity is %d\n", len(slstruct), cap(slstruct))

	//make function is used to initlialize and allocate memory for DS like slice,map and channels
	//make(type,size,capacity)

	s := make([]int, 10)
	fmt.Println(s)
}
