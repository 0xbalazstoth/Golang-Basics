package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

type desPerson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

type user struct {
	First string
	Age   int
}

type desUser struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	// DeserializeJSON()
	// handsOnExercise2()
	handsOnExercise4()
}

func SerializeJSON() string {
	p1 := person{
		First: "Bob",
		Last:  "Smith",
		Age:   27,
	}

	p2 := person{
		First: "Boby",
		Last:  "Smithy",
		Age:   22,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	return string(bs)
}

func DeserializeJSON() {
	jsonData := SerializeJSON()
	bs := []byte(jsonData)
	people := []desPerson{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range people {
		fmt.Println(i, v.First, v.Last, v.Age)
	}
}

func handsOnExercise1() string {
	u1 := user{
		First: "Bob",
		Age:   27,
	}

	u2 := user{
		First: "Boba",
		Age:   20,
	}

	u3 := user{
		First: "Boby",
		Age:   22,
	}

	users := []user{u1, u2, u3}

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range users {
		fmt.Println(i)
		fmt.Printf("Name: %v\nAge: %v\n", v.First, v.Age)
	}
	return string(bs)
}

func handsOnExercise2() {
	ds := handsOnExercise1()
	bs := []byte(ds)
	users := []desUser{}
	err := json.Unmarshal(bs, &users)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range users {
		fmt.Println(i)
		fmt.Printf("First: %v\nAge: %v\n", v.First, v.Age)
	}
}

func handsOnExercise3() {
	u1 := user{
		First: "Bob",
		Age:   27,
	}

	u2 := user{
		First: "Boba",
		Age:   20,
	}

	u3 := user{
		First: "Boby",
		Age:   22,
	}
	users := []user{u1, u2, u3}
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}

func handsOnExercise4() {
	xi := []int{5, 43, 23, 45, 23, 2465, 245, 1234, 24, 64, 57, 345, 235, 3, 46}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry"}

	for i := 0; i < len(xi)-1; i++ {
		if xi[i] > xi[i+1] {
			temp := xi[i]
			xi[i] = xi[i+1]
			xi[i+1] = temp
			i = -1
		}
	}
	fmt.Println(xi)

	for i := 0; i < len(xs)-1; i++ {
		if xs[i] > xs[i+1] {
			temp := xs[i]
			xs[i] = xs[i+1]
			xs[i+1] = temp
			i = -1
		}
	}
	fmt.Println(xs)
}
