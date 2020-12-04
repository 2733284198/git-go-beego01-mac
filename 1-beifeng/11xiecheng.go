package main

import (
	"fmt"
	"time"
)

func testRoutine() {
	fmt.Println("this is Routine=这是北风网go Runtine")
}

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func testfor() {
	for i := 1; i < 10; i++ {
		go Add(i, i)
	}
}

func main() {

	for i := 1; i < 10; i++ {
		go Add(i, i)
	}

	//testfor()
	time.Sleep(2)
}

func test() {
	go testRoutine()

	time.Sleep(2)
}

/*

beifeng #gosetup
2
6
8
12
18
14
10
16
4


*/
