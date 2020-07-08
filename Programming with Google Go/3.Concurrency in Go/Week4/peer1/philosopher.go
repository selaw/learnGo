package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Number of philosophers is simply the length of this list.
var ph = []string{"Philosopher1", "Philosopher2", "Philosopher3", "Philosopher4", "Philosopher5"}

const hunger = 3               // Number of times each philosopher eats
const wait = time.Second / 100 // Mean wait time
const eat = time.Second / 100  // Mean eat time

var fmt = log.New(os.Stdout, "", 0)

var dining sync.WaitGroup

func diningProblem(phName string, dominantHand, otherHand *sync.Mutex) {
	fmt.Println(phName, "Seated and waiting")
	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		fmt.Println(phName, "Eating")
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		fmt.Println(phName, "finished eating and waiting")
		rSleep(wait)
	}
	fmt.Println(phName, "Satisfied")
	dining.Done()
	fmt.Println(phName, "Left the table")
}

func main() {
	fmt.Println("Table empty")
	fmt.Println("dining begins")
	dining.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go diningProblem(ph[0], fork0, forkLeft)
	dining.Wait() // wait for philosphers to finish
	fmt.Println("Table empty")
	fmt.Println("dining end")
}
