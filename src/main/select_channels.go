package main

import "fmt"

func product(v1 int, v2 int, c chan int) {
	c <- v1 * v2
}

func main() {
	c := make(chan int)
	go product(1, 2, c)
	go product(3, 4, c)
	select {
	case a := <-c:
		fmt.Println(a)
	case b := <-c:
		fmt.Println(b)
	}
}
