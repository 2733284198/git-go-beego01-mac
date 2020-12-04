package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n\n ====> 14 Select <==== ")

	ch := make(chan int)

	select {
	case <-ch:
		fmt.Println("1-开始读")
	case ch <- 1:
		fmt.Println("2-开始写")
	default:
		fmt.Println("3-default")
	}

}
