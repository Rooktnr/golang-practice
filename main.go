package main

import "fmt"

func main() {

	// fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	// fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)

	//for loop
	// initialize; conditional; increment
	for i := 60; i < 123; i += 1 {
		fmt.Printf("%d \t %b \t %#x \t %q \n ", i, i, i, i)
	}

}
