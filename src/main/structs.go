package main

type PersonOld struct {
	name  string
	addr  string
	phone string
}

func main() {
	p1 := new(PersonOld)
	p1.name = "小明"
	p1.addr = "北今"
	p1.phone = "13888888888"

	println(p1)
	println(p1.name)
}
