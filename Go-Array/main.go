package main

import (
	"fmt"
)

func main() {
	//x := type{values}
	// ArrayNSlice()
	// Range()
	// SlicingSlice()
	// Append()
	// Delete()
	// Make()
	// handsOnExercise1()
	// handsOnExercise2n3()
	// handsOnExercise4()
	// handsOnExercise5()
	// handsOnExercise6()
	// handsOnExercise7()
	handsOnExercise8()
}

func ArrayNSlice() {
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

}

func Range() {
	x := []int{4, 5, 6, 7, 8}
	for i, v := range x {
		fmt.Println(i, v)
	}
}

func SlicingSlice() {
	x := []int{4, 45, 34, 54}
	fmt.Println(x[1:2])
}

func Append() {
	x := []int{4, 23, 45, 3}
	x = append(x, 43, 23, 21)
	fmt.Println(x)
}

func Delete() {
	x := []int{43, 23, 32, 23, 32}
	x = append(x[1:1], x[1:1]...)
	fmt.Println(x)
}

func Make() {
	x := make([]int, 10, 100)
	fmt.Println(x)
}

func handsOnExercise1() {
	arr := [5]int{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Printf("\n%T", arr)
}

func handsOnExercise2n3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

func handsOnExercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func handsOnExercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func handsOnExercise6() {
	x := []string{"Alabama", "Alaska", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia"}
	fmt.Printf("Length of slice: %v\n", len(x))
	fmt.Printf("Capacity of slice: %v\n", cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Printf("Index: %v; Value: %v\n", i, x[i])
	}
}

func handsOnExercise7() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Hellooo, James"}

	xxs := [][]string{xs1, xs2}

	for _, xs := range xxs {
		for j, v := range xs {
			fmt.Printf("%v %v\n", j, v)
		}
	}
}

func handsOnExercise8() {
	m := map[string][]string{
		"key":  []string{"value1", "value2"},
		"key2": []string{"value1", "value2"},
	}

	m["key3"] = []string{"value1", "value3"}
	delete(m, `key2`)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
