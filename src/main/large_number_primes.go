package main

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

func worker(id int, jobs <-chan *big.Int, results chan<- *big.Int, rounds int) {
	for num := range jobs {
		if num.ProbablyPrime(rounds) {
			results <- num
			fmt.Printf("Worker %d found prime: %v\n", id, num)
		}
	}
}

func main() {
	// Get the starting number from user input
	var start string
	fmt.Print("Enter the starting number: ")
	fmt.Scanln(&start)

	// Convert the starting number to a big.Int
	num := new(big.Int)
	num, _ = num.SetString(start, 10)

	// Get the number of threads from user input
	var numThreadsStr string
	fmt.Print("Enter the number of threads: ")
	fmt.Scanln(&numThreadsStr)

	// Convert the number of threads to an int
	numThreads, _ := strconv.Atoi(numThreadsStr)

	// Get the number of rounds from user input
	var roundsStr string
	fmt.Print("Enter the number of rounds: ")
	fmt.Scanln(&roundsStr)

	// Convert the number of rounds to an int
	rounds, _ := strconv.Atoi(roundsStr)

	// Create the jobs channel and the results channel
	jobs := make(chan *big.Int, numThreads*2)
	results := make(chan *big.Int, numThreads)

	// Start the worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results, rounds)
		}(i)
	}

	// Add jobs to the jobs channel
	for {
		// Check if the current number is prime
		if num.ProbablyPrime(rounds) {
			fmt.Printf("Found prime: %v\n", num)
			results <- num
		}

		// Increment the current number
		num = num.Add(num, big.NewInt(1))

		// Send the current number to the jobs channel
		jobs <- new(big.Int).Set(num)

		// Check if the results channel has any primes
		select {
		case prime := <-results:
			fmt.Printf("Found prime: %v\n", prime)
		default:
		}
	}

	// Wait for the worker goroutines to finish
	close(jobs)
	wg.Wait()
}
