package main

import "fmt"

//
//实现strStr()函数。
//
//给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。

//说明：
//
//当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
//对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。

//示例 1：
//
//输入：haystack = "hello", needle = "ll"
//输出：2

/////////////////暴力算法/////////////////
func strStr(haystack string, needle string) int {
	hashmap := []byte{}
	for j := 0; j < len(haystack); j++ {
		hashmap = append(hashmap, haystack[j])
	}
	fmt.Println(hashmap)
	for n := 0; n < len(haystack)-len(needle); n++ {

		if needle[0] == haystack[n] {
			fmt.Println("第一个相同的元素：", haystack[n])
			for m := 0; m < len(needle); m++ {
				if needle[m] == haystack[n] {
					fmt.Println(needle[m], haystack[n])
					n++
				}
				return -1
			}
		}
	}
	return 0
}

////////////////////////////////////////////////
func strStr2(haystack, needle string) int {
	n := len(haystack)
	m := len(needle)
outer:
	for i := 0; i <= n-m; i++ {
		for j := range needle {
			fmt.Println(haystack[i+j], needle[j])
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr2(haystack, needle))
}
