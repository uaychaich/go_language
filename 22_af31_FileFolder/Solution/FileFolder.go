package main

import (
	"fmt"
	"os"
)

func main() {
	curdir, err := os.Getwd()
	if err == nil {
		fmt.Printf("Current Working Directory: %v\n", curdir)
	}
	///////////////////////////////////////////////////////////////
	// os.Mkdir("./UayFolder", os.ModePerm)
	// os.Create("./UayFolder/data.txt")
	// ///////////////////////////////////////////////////////////////
	// file, err := os.OpenFile("./UayFolder/data.txt", os.O_RDWR, os.ModeDir)
	// if err == nil {
	// 	fmt.Fprintf(file, "Write:%v\n", time.Now().String())
	// }
	// file.Close()
	///////////////////////////////////////////////////////////////
	// b := []byte{}
	// b, err = os.ReadFile("./UayFolder/data.txt")
	// fmt.Printf("Content:%v", string(b))
	///////////////////////////////////////////////////////////////
	// os.RemoveAll("./UayFolder/")
}
