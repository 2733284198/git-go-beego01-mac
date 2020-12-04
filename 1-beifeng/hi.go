package main

//var x = 10 // error

import "fmt"

var x = 10

//x := 10 // syntax error: non-declaration statement outside function body

func main() {
	fmt.Println(x)
}
