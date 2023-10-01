package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x))
	}
	time.Sleep(1 * time.Second)
	printSomething("This is the second thing to be printed")
}