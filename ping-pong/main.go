package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <- ping
		if !ok {
			// do something
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
		fmt.Println("Executing loop")
	}
}

func main() {

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press Enter (enter Q to quit)")
	
	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput
		// wait for response
		response := <- pong
		fmt.Println(response)
	}
	close(ping)
	close(pong)
	fmt.Println("End of the program")
}