package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// go func(s string) {
	// 	for i := 0; i < 5; i++ {
	// 		time.Sleep(100 * time.Millisecond)
	// 		fmt.Println(s)
	// 	}
	// }("Uay")
	// go func1("Naja")
	// func1("Kiki")
	//////////////////////////////////////////////////////////////

	var wg sync.WaitGroup
	wg.Add(2)
	go func2("Uay", &wg)
	go func2("Kiki", &wg)
	wg.Wait()
}

func func1(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func func2(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
