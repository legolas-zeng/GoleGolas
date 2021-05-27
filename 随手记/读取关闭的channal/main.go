package main

import (
	"fmt"
)

/////////////////有缓冲//////////////////

//func main(){
//	wg := sync.WaitGroup{}
//	ch1 := make(chan int,5)
//	wg.Add(1)
//	go func() {
//		fmt.Println("存放信息")
//		for i:=0;i<5;i++{
//			ch1<-i
//		}
//		close(ch1)
//		fmt.Println("关闭channal")
//		wg.Done()
//	}()
//	wg.Wait()
//	fmt.Println("read")
//	for i := 1; i < 9; i++ {
//		v, ok := <-ch1
//		fmt.Println(v)
//		fmt.Println(ok)
//	}
//}

////////////////无缓存//////////////////////////

func main() {
	ch1 := make(chan int)
	go func() {
		fmt.Println("存放信息")
		ch1 <- 1
		close(ch1)
		fmt.Println("关闭channal")
	}()
	fmt.Println("read")
	for i := 1; i < 5; i++ {
		v, ok := <-ch1
		fmt.Println(v)
		fmt.Println(ok)
	}
}
