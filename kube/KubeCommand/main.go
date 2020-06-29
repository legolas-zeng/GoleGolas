//Author: zwa
//Date: 2020/6/29

package main

import (
	"flag"
	"fmt"
)

func main() {
	var address string
	flag.StringVar(&address, "address", "GuangZhou", "Where is your address?")
	fmt.Println(address)
}
