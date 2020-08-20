package main

import "fmt"

func main() {
	// _Loops1()
	// _Loops2()
	// _Loops3()
	// _Loops4()
	_Loops5()
}

func _Loops1() {
	for i := 1; i < 10000+1; i++ {
		fmt.Println(i)
	}
}

func _Loops2() {
	for i := 65; i < 90+1; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}

	}
}

func _Loops3() {
	bd := 1991
	for bd <= 2017 {
		fmt.Println(bd)
		bd++
	}
}

func _Loops4() {
	bd := 1991
	for {
		if bd > 2017 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

func _Loops5() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Printf("%v %v\n", i, i%4)
		}
	}
}
