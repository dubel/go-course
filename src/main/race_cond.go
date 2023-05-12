/*
 * This program creates 2 go routines which share the variable v, one go routine
 * is to increment this integer and the other goroutine to print its value. Since
 * there is no synchronisation mechanism between the two threads and the data is
 * shared, it is most likely to print different value after each run. The value
 * printed cannot be predicted, hence the race condition.
 */

package main

import (
	"fmt"
	"time"
)

var v int = 0

func main() {

	go change()
	go print()

	time.Sleep(time.Second)
	fmt.Println("Done!")
}

func print() {
	fmt.Println("The value of v is", v)
}

func change() {
	for i := 0; i < 1000000; i++ {
		v = i
	}
}
