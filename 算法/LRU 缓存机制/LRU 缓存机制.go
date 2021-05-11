package main

import "fmt"

type LRUnode struct {
	key int
	value int
	pre *LRUnode
	next *LRUnode
}

var LRUCache map[int] *LRU

func Constructor(){
	cache :=
}

func main(){
	LRUCache = map[int] *LRU{}
	LRUCache[1]=&LRU{1,1}
	LRUCache[2]=&LRU{2,2}
	fmt.Println(LRUCache)
}
