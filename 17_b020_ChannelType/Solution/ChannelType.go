package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// var ch1 chan int
	// ch1 = make(chan int, 2)
	// ch1 <- 5
	// ch1 <- 7
	// fmt.Printf("len(ch1)=%v\n", len(ch1))
	// res1 := <-ch1
	// fmt.Printf("len(ch1)=%v | res1=%v\n", len(ch1), res1)
	// fmt.Printf("//////////////////////////////////////////////\n")
	//////////////////////////////////////////////////////////////////
	// ch2 := make(chan int, 2)
	// ch2 <- 5
	// ch2 <- 7
	// close(ch2)
	// for x := range ch2 {
	// 	fmt.Printf("x=%v | len(ch2)=%v\n", x, len(ch2))
	// }
	// ch2 = make(chan int, 2)
	// ch2 <- 8
	// close(ch2)
	// for x := range ch2 {
	// 	fmt.Printf("x=%v | len(ch2)=%v\n", x, len(ch2))
	// }
	// fmt.Printf("//////////////////////////////////////////////\n")
	//////////////////////////////////////////////////////////////////
	// var sw sync.WaitGroup
	// var ch1 chan string = make(chan string)
	// sw.Add(1)
	// go Uayfunc1(ch1, &sw)
	// time.Sleep(5 * time.Second)
	// ch1 <- "Hello Uaychai"
	// sw.Wait()
	// // ch1 <- "Hello Uaychai"
	// sw.Add(1)
	// go Uayfunc1(ch1, &sw)
	// time.Sleep(5 * time.Second)
	// ch1 <- "Hello Uaychai Again"
	// sw.Wait()
	// fmt.Printf("//////////////////////////////////////////////\n")
	////////////////////////////////////////////////////////////////
	// var sw sync.WaitGroup
	// var ch1 chan string = make(chan string)
	// var ch2 chan int = make(chan int)
	// sw.Add(1)
	// go Uayfunc2(ch1, ch2, &sw)
	// time.Sleep(5 * time.Second)
	// ch1 <- "Hello Uaychai"
	// sw.Wait()
	// sw.Add(1)
	// go Uayfunc2(ch1, ch2, &sw)
	// time.Sleep(5 * time.Second)
	// ch2 <- 500
	// sw.Wait()
	// fmt.Printf("//////////////////////////////////////////////\n")
	////////////////////////////////////////////////////////////////
	var sw sync.WaitGroup
	var ch1 chan string = make(chan string)
	var qch chan int = make(chan int)
	sw.Add(1)

	go Uayfunc4(ch1, qch, &sw)
	time.Sleep(5 * time.Second)
	fmt.Printf("%v\n", <-ch1)
	fmt.Printf("%v\n", <-ch1)
	qch <- 1
	sw.Wait()
	fmt.Printf("//////////////////////////////////////////////\n")
	////////////////////////////////////////////////////////////////
}
