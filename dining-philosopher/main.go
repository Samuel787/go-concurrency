package main

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

// Philosopher is a struct which stores information about a philosopher.
type Philosopher struct {
	name string
	rightFork int
	leftFork int
}

// philosophers is list of all philosophers.
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define some variables
var hunger = 3 // how many times does a person eat?
var eatTime = 100 * time.Millisecond
var thinkTime = 300 * time.Millisecond
var sleepTime = 100 * time.Millisecond

var philosophersOrder = []string{}

func main() {
	// print out a welcome message
	color.Cyan("Dining Philosopher Problem")
	color.Cyan("--------------------------")
	color.Cyan("The table is empty.")

	// start the meal
	dine()

	// print out finished message
	color.Cyan("The table is empty")
}

func dine() {
	wg := &sync.WaitGroup{} // wait until done eating
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{} // wait until everyone is seated
	seated.Add(5)

	orderLock := &sync.Mutex{}

	// forks is a map of all 5 forks.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated, orderLock)
	}

	wg.Wait()

	var ordering = ""
	for _, name := range philosophersOrder {
		ordering += name + " "
	}
	color.Cyan("This is the order of eating: %s", ordering)
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup, orderLock *sync.Mutex) {
	defer wg.Done()

	// seat the philosopher at the table
	color.Blue("%s is seated at the table.", philosopher.name)
	seated.Done()

	seated.Wait()

	// eat 3 times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			color.Green("%s takes the right fork.", philosopher.name)
			forks[philosopher.leftFork].Lock()
			color.Green("%s takes the left fork.", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			color.Green("%s takes the left fork.", philosopher.name)
			forks[philosopher.rightFork].Lock()
			color.Green("%s takes the right fork.", philosopher.name)
		}

		color.Green("%s has both forks and is eating.", philosopher.name)
		time.Sleep(eatTime)
		
		color.Green("%s is thinking.", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		color.Green("%s put down the forks.", philosopher.name)
	}

	orderLock.Lock()
	philosophersOrder = append(philosophersOrder, philosopher.name)
	orderLock.Unlock()
	
	color.Blue("%s is satisfied", philosopher.name)
	color.Blue("%s has left the table.", philosopher.name)
}