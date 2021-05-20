package main

import (
	"fmt"
	"sync"
)

//////////////////select随机性/////////////////////////
//func main() {
//	runtime.GOMAXPROCS(1)
//	int_chan := make(chan int, 1)
//	string_chan := make(chan string, 1)
//	int_chan <- 1
//	string_chan <- "hello"
//	select {
//	case value := <-int_chan:
//		fmt.Println(value)
//	case value := <-string_chan:
//		panic(value)
//	}
//}
//select会随机选择一个可用通用做收发操作,所以这个有可能触发异常
/////////////////////////////////////////////////

//func cat(ch chan string) {
//	for i := 0; i < 100; i++ {
//		ch <- "cat"
//	}
//	close(ch)
//
//}
//
//func dog(ch chan string) {
//	for i := 0; i < 100; i++ {
//		ch <- "dog"
//	}
//	close(ch)
//}
//
//func main() {
//	ch1 := make(chan string)
//	ch2 := make(chan string)
//	//ch3 := make(chan string)
//	go cat(ch1)
//	go dog(ch2)
//	//go fish(ch3)
//	for a := range ch1 {
//		fmt.Println(a)
//	}
//	for b := range ch2 {
//		fmt.Println(b)
//	}
//}

//////////////channal通信/////////////////
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	fmt.Println("main begin")
//	go func() {
//		for {
//			ch1 <- 1
//			<-ch2
//			fmt.Println("你好")
//		}
//	}()
//	go func() {
//		for {
//			<-ch1
//			ch2 <- 1
//			fmt.Println("小仙女")
//		}
//	}()
//
//	fmt.Println("main end")
//	time.Sleep(1 * time.Second)
//}

//////////////////// 循环/////////////////////
var wg sync.WaitGroup

func do1(ch1 chan int, ch2 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- 1
		<-ch2
		fmt.Println("你好")
	}
	close(ch1)
	close(ch2)
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	fmt.Println("main begin")
	wg.Add(10)
	wg.Add(1)
	go do1(ch1, ch2)
	wg.Done()
	go func() {
		for i := 0; i < 10; i++ {
			<-ch1
			ch2 <- 1
			fmt.Println("小仙女,收到")
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println("main close")
}
