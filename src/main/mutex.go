package main

import (
	"fmt"
	"sync"
)

var i int = 0
var mut sync.Mutex

func inc(group *sync.WaitGroup) {
	mut.Lock()
	i++
	mut.Unlock()
	group.Done()
}

func main() {
	group := sync.WaitGroup{}
	group.Add(2)
	go inc(&group)
	go inc(&group)
	group.Wait()
	fmt.Println(i)
}
