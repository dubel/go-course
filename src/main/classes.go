package main

type Point struct {
	X, Y int
}

type MyInt int

// it's a receiver function
// https://medium.com/@adityaa803/wth-is-a-go-receiver-function-84da20653ca2
// myInt is passed by value implicitly
func (myInt MyInt) DoubleValue() int {
	return int(myInt * 2)
}

func main() {
	var myInt MyInt = 5
	println(myInt.DoubleValue())
}
