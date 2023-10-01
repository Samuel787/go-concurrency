package main

import (
	"io"
	"os"
	"strings"
	"testing"
	"sync"
	"fmt"
)

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("epsilon", &wg)
	wg.Wait()

	if msg != "epsilon" {
		t.Errorf("Expected epsilon but got %s", msg)
	}
}


func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "epsilon"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	
	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon but found %s", output)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	fmt.Printf("This is the output %s: ", output)

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected to find Hello, universe!")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Expected to find Hello, cosmos!")
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected to find Hello, world!")
	}

}