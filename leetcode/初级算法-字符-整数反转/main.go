package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x < 0 {
		fmt.Println("运算负数")
		return -reverse(-x)
	}
	if x > math.MaxInt32 || x < math.MinInt32 {
		fmt.Println("超出范围")
		return 0
	}
	var tmp int32
	for x > 0 {
		//fmt.Println("个位数：",x%10)
		var d int32 = int32(x % 10)
		x = x / 10 //高位数
		if 10*tmp/10 != tmp {
			return 0
		}
		tmp = 10*tmp + d
		//fmt.Println(tmp)
	}
	return int(tmp)
}

func main() {
	x := 1534236469
	fmt.Println(reverse(x))
}
