package main

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex

func add() {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Println("加法", count)
}

func substruct() {
	lock.Lock()
	defer lock.Unlock()
	count--
	fmt.Println("减法", count)

}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
}
