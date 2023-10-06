package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i ++ {
		philosophersOrder = []string{}
		dine()
		if len(philosophersOrder) != 5 {
			t.Errorf("Incorrect length of slice; expected 5 but got %d", len(philosophersOrder))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct{
		name string
		delay time.Duration
	} {
		{"zero second delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		philosophersOrder = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		dine()
		if len(philosophersOrder) != 5 {
			t.Errorf("Incorrect length of slice; expected 5 but got %d for delay: %s", len(philosophersOrder), e.name)
		}
	}
}

