package main

import (
	"fmt"
	"sync"
	"time"
)

type SyncMap struct {
	mymap         map[string]string
	*sync.RWMutex //读写锁
}

//var smap SyncMap //公有的访问map
var done chan bool //通道，是否完成

func main() {
	smap := new(SyncMap)
	smap.mymap = map[string]string{}
	smap.RWMutex = new(sync.RWMutex)
	done = make(chan bool, 1000)

	go func() {
		for {
			smap.Lock()
			smap.mymap["1"] = "1"
			smap.Unlock()
			done <- true
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func() {
		for {
			smap.Lock()
			smap.mymap["1"] = "2"
			smap.Unlock()
			done <- true
			time.Sleep(1 * time.Millisecond)
		}
	}()

	var lastlength = 0
	var lock sync.Mutex
	go func() {
		for {
			if len(done) != lastlength {
				lock.Lock()
				lastlength = len(done)
				lock.Unlock()

				smap.RLock()
				fmt.Print(smap.mymap["1"], "\t")
				if len(done)%5 == 0 {
					fmt.Println("  ", lastlength)
				}
				smap.RUnlock()
			}
		}
	}()

	for {
		if len(done) == 1000 {
			fmt.Println("通道已经满了")
			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
