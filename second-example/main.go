package main

import (
	"sync"
	"fmt"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	msg = s
	mutex.Unlock()
}


func main() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe", &mutex)
	go updateMessage("Hello, cosmos", &mutex)
	wg.Wait()

	fmt.Println(msg)
}