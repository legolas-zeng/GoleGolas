package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func A() {
	defer wg.Done()
	fmt.Println("A")
}

func B() {
	defer wg.Done()
	fmt.Println("B")
}

func main() {
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
}
