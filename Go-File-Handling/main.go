package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First string
	Last  string
	Age   int
}

type readPerson struct {
	First string
	Last  string
	Age   string
}

func main() {
	WriteNames()

	for _, v := range ReadNames() {
		fmt.Printf("Firstname: %v\nLastname: %v\nAge: %v\n", v.First, v.Last, v.Age)
	}
}

func WriteNames() {
	p1 := person{
		First: "Bob",
		Last:  "S.",
		Age:   20,
	}

	p2 := person{
		First: "Boby",
		Last:  "TS.",
		Age:   22,
	}

	p3 := person{
		First: "Boba",
		Last:  "SS.",
		Age:   23,
	}

	people := []person{p1, p2, p3}

	file, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Sync()
	defer file.Close()

	for i := 0; i < len(people); i++ {
		fmt.Fprintf(file, "First: %v,Last: %v,Age: %v\n", people[i].First, people[i].Last, people[i].Age)
	}
}

func ReadNames() []readPerson {
	read, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)
	var arr []string

	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	read.Close()

	xArr := []readPerson{}

	for _, val := range arr {
		split := strings.Split(val, ",")
		firstName := strings.Replace(split[0], "First: ", "", -1)
		lastName := strings.Replace(split[1], "Last: ", "", -1)
		age := strings.Replace(split[2], "Age: ", "", -1)

		n := readPerson{First: firstName, Last: lastName, Age: age}
		xArr = append(xArr, n)
	}

	return xArr
}
