package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, CSticks[i], CSticks[(i+1)%5], 0}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat()
	}
	wg.Wait()
}

type ChopS struct{ sync.Mutex }

type Philosopher struct {
	number          int
	leftCS, rightCS *ChopS
	eatCount        int
}

var wait sync.Mutex

func (p Philosopher) eat() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		wait.Lock()
		if (p.eatCount) < 3 {
			rand := rand.Intn(2)
			// let picking chopsticks in random order
			switch rand {
			case 0:
				p.leftCS.Lock()
				p.rightCS.Lock()
			case 1:
				p.rightCS.Lock()
				p.leftCS.Lock()
			}
			wait.Unlock()
			p.eatCount = p.eatCount + 1
			fmt.Println("starting to eat ", p.number, " [for the ", p.eatCount, " time]")
			time.Sleep(time.Second) // Simulate eating
			fmt.Println("finishing eating ", p.number)

			switch rand {
			case 0:
				p.leftCS.Unlock()
				p.rightCS.Unlock()
			case 1:
				p.rightCS.Unlock()
				p.leftCS.Unlock()
			}
		} else {
			wait.Unlock()
		}
	}
}
