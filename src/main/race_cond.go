package main

/* Two methods add and sub - are both accessing shared state (of an integer) incrementing and decrementing it - respectively.
Execution of those method is done using goroutines using keyword "go" - leading to race condition.
Most of the time the output will be "1" but sometimes it will be "0" depending on an outcome of interleaves.
Program executes infinitely - until sigkill (ctrl/cmd + C) is sent.

*/
import (
	"fmt"
	"time"
)

var x int

func add(a *int) {
	*a += 1
}

func sub(a *int) {
	*a -= 1
}

func main() {

	for {
		x := 1
		go sub(&x)
		go add(&x)
		fmt.Print(x)
		time.Sleep(1 * time.Second)
	}
}
