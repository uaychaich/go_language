package main

import (
	"fmt"
	"sync"
	"time"
)

func Uayfunc1(ch chan string, sw *sync.WaitGroup) {
	defer sw.Done()
	fmt.Printf("%v\n", <-ch)
}

func Uayfunc2(datach1 chan string, datach2 chan int, sw *sync.WaitGroup) {
	defer sw.Done()
	select {
	case data1 := <-datach1:
		fmt.Printf("External send string to datach1:%v\n", data1)
	case data2, ok := <-datach2:
		if ok {
			fmt.Printf("External send int to datach2:%v\n", data2)
		}
		// default:
		// 	fmt.Printf("I don't know")
	}
}

func Uayfunc3(datach1 chan string, sw *sync.WaitGroup) {
	defer sw.Done()
	data1 := fmt.Sprintf("Uaychai:%v", time.Now().String())
	select {
	case datach1 <- data1:
		fmt.Printf("External receive string from datach1:%v\n", data1)
		// default:
		// 	fmt.Printf("I don't know")
	}
}

func Uayfunc4(datach1 chan string, quitch chan int, sw *sync.WaitGroup) {
	defer sw.Done()
	for {
		data1 := fmt.Sprintf("Uaychai:%v", time.Now().String())
		select {
		case datach1 <- data1:
			fmt.Printf("External receive string from datach1:%v\n", data1)
		case <-quitch:
			return
			// default:
			// 	fmt.Printf("I don't know")
		}
	}
}
