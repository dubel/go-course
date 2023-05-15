package main

import (
	"fmt"
	"sync"
)

func makeValue(c chan int, wg *sync.WaitGroup) {
	i := 0
	for {
		if i > 10000 {
			close(c)
		}
		c <- i
		i++
	}

}

func readChannel(c chan int, wg *sync.WaitGroup) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	go makeValue(c, &wg)
	go readChannel(c, &wg)

	wg.Wait()
	fmt.Println("Maine routine")
}
