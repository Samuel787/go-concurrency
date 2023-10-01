package main

import (
	"sync"
	"testing"
)
// run this with $ go test -race .
func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(1)
	go updateMessage("Goodbye, cruel world!", &mutex)
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Errorf("incorrect value in msg")
	}
}