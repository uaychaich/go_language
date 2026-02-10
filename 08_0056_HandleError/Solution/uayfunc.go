package main

import "log"

func UayFunc1(x int, y int) (a int) {
	a = x / y
	return
}

func UayFunc2(x int, y int) (a int) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	a = x / y
	return
}

func UayFunc3(year uint16) (age uint16) {
	if year >= 1000 && year <= 9999 {
		age = 2026 - year
	} else {
		panic("year is incorrect")
	}
	return
}
