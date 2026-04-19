package main

import (
	"fmt"
)

func mayPanic() {
	fmt.Println("Before panic")
	panic("Something bad happened")
	// fmt.Println("After panic") // Никогда не будет выполнен
}

func handlePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Caught panic:", r)
		}
	}()

	mayPanic()
	fmt.Println("Continuing execution after recovering from panic")
}

func main() {
	handlePanic()
	fmt.Println("Program completes successfully")
}
