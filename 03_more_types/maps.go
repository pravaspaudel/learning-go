package main

import "fmt"

//map are key value pairs. its zero value is nil
var m map[string]int

type Circle struct {
	r float32
}

func maps() {
	fmt.Println("maps in go")

	m := make(map[string]int)

	m["a"] = 32
	m["b"] = 30
	m["c"] = 34
	m["d"] = 39
	m["e"] = 50

	// for key, val := range m {
	// 	fmt.Println(key, val)
	// }
	var m1 = map[string]Circle{
		"bigcircle":    {r: 40.45},
		"mediumcircle": {r: 30.45},
		"smallcircle":  {r: 20.45},
	}
	fmt.Println(m1)
	fmt.Println(m1["bigcircle"])

	//operations on map
	//insert or update or delete an element in map
	m1["bigcircle"] = Circle{r: 43}
	delete(m1, "bigcircle")

	//insert
	m1["circle"] = Circle{r: 60}
	fmt.Println(m1)

	ele, ok := m1["notkey"]
	if ok {
		fmt.Println("keypresent", ele)
	} else {
		fmt.Println("key not present")
	}
}
