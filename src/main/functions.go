package main

func main() {
	println(foo(1, 2))
}

func foo(x, y int) (int, int) {
	return x + y, x - y
}
