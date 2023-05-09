package main

import (
	"fmt"
)

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
}

func sieve(ch chan int) {
	prime := <-ch
	fmt.Println(prime)

	ch1 := make(chan int)
	go filter(ch, ch1, prime)
	sieve(ch1)
}

func main() {
	var numThreads int
	fmt.Print("Enter the number of threads: ")
	fmt.Scan(&numThreads)

	ch := make(chan int)
	go generate(ch)
	for i := 0; i < numThreads; i++ {
		go sieve(ch)
	}

	// Wait for user to press enter
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
