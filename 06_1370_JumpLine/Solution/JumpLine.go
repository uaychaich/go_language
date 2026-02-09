package main

import (
	"fmt"
)

func main() {
	goto des1
	fmt.Println("Uaychai naja")
des1:
	fmt.Println("des1 Line naja")
	////////////////////////////////////////////////////////////
	for i2 := 0; i2 < 3; i2++ {
		for j2 := 0; j2 < 3; j2++ {
			if i2 == 0 && j2 == 1 {
				continue
			}
			if i2 == 1 && j2 == 1 {
				break
			}
			fmt.Printf("i2=%d | j2=%d\n", i2, j2)
		}
	}
	////////////////////////////////////////////////////////////
des2:
	for i3 := 0; i3 < 3; i3++ {
		for j3 := 0; j3 < 3; j3++ {
			if i3 == 0 && j3 == 1 {
				continue des2
			}
			if i3 == 1 && j3 == 1 {
				break des2
			}
			fmt.Printf("i3=%d | j3=%d\n", i3, j3)
		}
	}
	////////////////////////////////////////////////////////////
	for i3 := 0; i3 < 5; i3++ {
		switch i3 {
		case 0, 1, 2, 4:
			fmt.Printf("Normal: i3 is %d\n", i3)
		case 3:
			break
			fmt.Printf("Three")
		}
	}
	////////////////////////////////////////////////////////////
des3:
	for i4 := 0; i4 < 5; i4++ {
		switch i4 {
		case 0, 1, 2, 4:
			fmt.Printf("Normal: i4 is %d\n", i4)
		case 3:
			break des3
			fmt.Printf("Three")
		}
	}
	////////////////////////////////////////////////////////////
}
