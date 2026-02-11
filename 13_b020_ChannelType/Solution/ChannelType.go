package main

import "fmt"

func main() {
	var ch1 chan int
	ch1 = make(chan int, 2)
	ch1 <- 5
	ch1 <- 7
	fmt.Printf("len(ch1)=%v\n", len(ch1))
	res1 := <-ch1
	fmt.Printf("len(ch1)=%v | res1=%v\n", len(ch1), res1)
	fmt.Printf("//////////////////////////////////////////////\n")
	//////////////////////////////////////////////////////////////////
	ch2 := make(chan int, 2)
	ch2 <- 5
	ch2 <- 7
	close(ch2)
	for x := range ch2 {
		fmt.Printf("x=%v | len(ch2)=%v\n", x, len(ch2))
	}
	ch2 = make(chan int, 2)
	ch2 <- 8
	close(ch2)
	for x := range ch2 {
		fmt.Printf("x=%v | len(ch2)=%v\n", x, len(ch2))
	}
	fmt.Printf("//////////////////////////////////////////////\n")
	//////////////////////////////////////////////////////////////////

}
