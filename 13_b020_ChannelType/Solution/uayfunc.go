package main

import (
	"fmt"
	"sync"
)

func Uayfunc1(ch chan string, sw *sync.WaitGroup) {
	defer sw.Done()
	fmt.Printf("%v\n", <-ch)
}

func Uayfunc2(datach1 chan string, datach2 chan int, sw *sync.WaitGroup) {
	defer sw.Done()
	select {
	case data1 := <-datach1:
		fmt.Printf("Receive datach1:%v\n", data1)
	case data2 := <-datach2:
		fmt.Printf("Receive datach2:%v\n", data2)
		// default:
		// 	fmt.Printf("I don't know")
	}
}
