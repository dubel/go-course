package main

import (
	"fmt"
	"sync"
)

func subRoutine(wg *sync.WaitGroup) {
	fmt.Println("Done")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go subRoutine(&wg)
	wg.Wait()
	fmt.Println("Maine routine")
}
