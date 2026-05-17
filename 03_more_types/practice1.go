package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

// Create a Task struct {Title, Done}.
// Store tasks in a slice.
type Task struct {
	Title string
	Done  bool
}

var tasks = []Task{
	{"do task one", false},
	{"done task two", true},
	{"do task three", false},
	{"do task four", false},
	{"done task five", true},
	{"do task six", false},
}

// List only completed tasks.
func completedTasks(tasks []Task) []Task {
	var completed = []Task{}
	for _, task := range tasks {
		if task.Done == true {
			completed = append(completed, task)
		}
	}
	return completed
}

// count pending tasks
func pendingTasks(tasks []Task) int {
	count := 0
	for _, task := range tasks {
		if task.Done == false {
			count++
		}
	}
	return count
}

// Create a phonebook using map[string]string. Search for a contact.
// Handle case when contact doesn’t exist.
var phonebook = map[string]string{
	"abc":  "082084289520",
	"pqrs": "082084289521",
	"mnop": "182084289520",
	"yyxz": "082084289528",
}

// Create nested structs (User with Address struct).
type Address struct {
	city       string
	zipCode    string
	postalCode int64
}

type User struct {
	name    string
	age     int
	address Address
}

// Group words by length using a map.
func groupWordByLength(s string) {
	m := make(map[int][]string)
	words := strings.Split(s, " ")

	for _, word := range words {
		length := len(word)

		m[length] = append(m[length], word)
	}

	fmt.Println(m)
}

// Store multiple users in a map using ID as key.
var users = make(map[string]User)

func addUser(user User) {
	b := make([]byte, 8)
	_, err := rand.Read(b)

	if err != nil {
		panic(err)
	}
	s := hex.EncodeToString(b)
	// str := strings.Join(b, ",")
	fmt.Println("random Id : ", b, s)
	users[s] = user
}

// Create a slice of maps and simulate JSON-like data.
var data = []map[string]any{
	{"id": 1, "name": "xyz", "age": 32},
	{"id": 2, "name": "abc", "age": 30},
	{"id": 3, "name": "pqr", "age": 34},
	{"id": 4, "name": "mno", "age": 31},
}

func practice1() {
	fmt.Println("practice1.go")
	fmt.Println(pendingTasks(tasks))

	name := "notpresent"

	phone, ok := phonebook[name]

	if ok {
		fmt.Println("key does exists with val ", phone)
	} else {
		fmt.Println("key does not exists")
	}
	// Create a contact with mixed data like age, name
	var contact = map[string]interface{}{
		"name":   "abc",
		"age":    20,
		"active": false,
	}
	fmt.Println(contact)
	nam, ok := contact["name"]
	if ok {
		fmt.Println(nam, ok)
	} else {
		fmt.Println(nam, ok)
	}

	user := User{
		name: "abc",
		age:  45,
		address: Address{
			city:       "xyz",
			zipCode:    "zip",
			postalCode: 4000,
		},
	}

	user1 := User{
		name: "pqr",
		age:  4,
		address: Address{
			city:       "pqrcity",
			zipCode:    "zip",
			postalCode: 5000,
		},
	}

	addUser(user)
	addUser(user1)

	for key, val := range users {
		fmt.Println(key, val)
	}

	groupWordByLength("hello how are you hope you are doing well thank go hi")

}
