package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	a := "FC Barcelona"
	fmt.Println(&a)

	var b *string = &a
	fmt.Println(b)

	aa := 27
	var c *int = &aa
	// fmt.Println(c)
	// fmt.Println(&c) //Address of c variable
	// fmt.Println(*c) //Returns original value of c variable
	fmt.Println(*&c)

	val := 2700
	fmt.Println(&val)

	p := person{
		first: "Bob",
		last:  "S.",
	}

	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	p.first = "Bobby"
	p.last = "Smith"
}
