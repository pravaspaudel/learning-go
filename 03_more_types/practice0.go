package main

import (
	"fmt"
	"strings"
)

// Create a Student struct with Name, Age, Marks.
type Student struct {
	Name  string
	Age   int
	Marks float32
}

// Store 5 students in a slice and print them.
var Students = []Student{
	{"larry", 23, 90.5},
	{"marry", 20, 95.5},
	{"oarry", 18, 88.75},
	{"qarry", 19, 80},
	{"parry", 19, 83.5},
}

// Find the student with highest marks and average marks of all students
func calculateMax(students []Student) (Student, float32) {
	var avg float32 = 0
	maxStudent := students[0]
	for _, s := range students {
		avg += (s.Marks)
		if s.Marks > maxStudent.Marks {
			maxStudent = s
		}
	}
	avg = avg / float32(len(students))
	return maxStudent, avg
}

// Count word frequency from a sentence using a map
func countFrequency(sentence string) map[string]int {
	words := strings.Split(sentence, " ")
	freq := make(map[string]int)
	for _, ch := range words {
		freq[string(ch)]++
	}
	return freq
}

func practice0() {
	fmt.Println("practice0.go")
	fmt.Println(Students)
	fmt.Println(calculateMax(Students))

	//Create a map with key string and value of int for product inventory.
	inventory := make(map[string]int)
	// Add new items to the map.
	// Update quantity of an existing item.
	// Delete an item from inventory.
	inventory["apple"] = 34
	inventory["samsung"] = 37
	delete(inventory, "apple")
	fmt.Println(inventory)

	sentence := "hello hi , I am a human  hello ! a !"
	fmt.Println(countFrequency(sentence))

	//Create a slice of integers and remove all odd numbers.
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evens := []int{}

	for _, val := range nums {
		if val%2 == 0 {
			evens = append(evens, val)
		}
	}

	fmt.Println(evens)

}
