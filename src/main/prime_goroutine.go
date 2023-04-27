package main

import "fmt"

// Generate the input values passing through the go channel pipe
func generate(ch chan int) {
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
}

// Copy the value from input channel to output channel
func filter(in chan int, out chan int, filter int) {
	for val := range in {
		if val%filter != 0 {
			out <- val // send value to the output pipe
		}
	}
}

func main() {
	ch := make(chan int) // create a new a channel
	generate(ch)

	for {
		prime := <-ch
		fmt.Println(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
