package main

import "fmt"

func main() {
	// handsOnExercise1()
	// handsOnExercise2()
	// handsOnExercise3()
	// handsOnExercise4()
}

func handsOnExercise1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func handsOnExercise2() {
	var x int
	var y string
	var z bool
	fmt.Println(x, y, z)
}

func handsOnExercise3() {
	var x int = 42
	var y string = "James Bond"
	var z bool = true

	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}

func handsOnExercise4() {
	type number32 int32
	var x number32
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	var y int = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
