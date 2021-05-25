package main

import (
	"fmt"
	"math"
)

/////////////////////////////////////////////////////
//我这种思路有问题，应该先确定数字字符的开始索引，而不是遇到数字才开始匹配判断

//func myAtoi(s string) int {
//	re := 0
//	sign := 1
//	for i:=0;i<len(s);i++{
//		if s[i]>='0'&&s[i]<='9'{
//			tmp,_:=strconv.Atoi(string(rune(s[i])))
//			re = re*10+tmp
//		}else if s[i]=='-'{
//			sign = -1
//		}else if s[i]==' '|| s[i]=='+'{
//			continue
//		}else if s[i]=='-+'||s[i]=='+-'{
//			break
//		}
//	}
//	fmt.Println(re)
//	if re*sign > math.MaxInt32{
//		return math.MaxInt32
//	}else if re*sign < math.MinInt32 {
//		return math.MinInt32
//	}
//	return re*sign
//}
//////////////////////////////////////////////////////////
//

func myAtoi2(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	//丢弃无用的前导空格,将i挪到有字符的开始前
	for i < n && s[i] == ' ' {
		i++
	}
	fmt.Println(i)
	//标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
			fmt.Println(i)
		} else if s[i] == '+' {
			sign = 1
			i++
			fmt.Println(i)
		}
	}
	fmt.Println(i, n)
	for i < n && s[i] >= '0' && s[i] <= '9' {
		fmt.Println(s[i])
		abs = 10*abs + int(s[i]-'0') //字节 byte '0' == 48
		fmt.Println("结果：", abs)
		if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}

func main() {
	s := " -+12"
	fmt.Println(myAtoi2(s))
}
