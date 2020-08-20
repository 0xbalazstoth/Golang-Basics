package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

type person struct {
	first string
}

type human interface {
	speak()
}

func main() {
	// fmt.Println(runtime.GOOS)
	// fmt.Println(runtime.GOARCH)
	// fmt.Println(runtime.NumCPU())
	// fmt.Println(runtime.NumGoroutine())

	// wg.Add(1)

	// go helloWorld()

	// fmt.Println(runtime.NumCPU())
	// fmt.Println(runtime.NumGoroutine())
	// wg.Wait()

	// RaceCondition()
	// handsOnExercise1()
	handsOnExercise2()
}

func helloWorld() {
	fmt.Println("Hello")
	wg.Done()
}

func RaceCondition() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Routines: ", runtime.NumGoroutine())
	counter := 0
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Routines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Routines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}

func handsOnExercise1() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello 1")
		fmt.Println("Routine 1: ", runtime.NumGoroutine())
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello 2")
		fmt.Println("Routine 2: ", runtime.NumGoroutine())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("exit...")
	fmt.Println("Routine exit: ", runtime.NumGoroutine())
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func handsOnExercise2() {
	p1 := person{
		first: "Bob",
	}
	saySomething(&p1)
	p1.speak()
}
