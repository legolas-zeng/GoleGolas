package main

import (
	"fmt"
)

/////////////////////////////////////////
// 先把s存入hash表中， 对比t，然后获得结果
func isAnagram(s string, t string) bool {
	hashmap := [26]int{}
	for _, v := range s {
		hashmap[v-'a']++
	}
	fmt.Println(hashmap)
	for _, v := range t {
		if hashmap[v-'a'] != 0 {
			hashmap[v-'a']--
		} else {
			hashmap[v-'a']++
		}
	}
	fmt.Println(hashmap)
	for _, v := range hashmap {
		if v != 0 {
			return false
		}
	}
	return true
}

/////////////////////////////////////////////
// 切片可以直接对比返回
func isAnagram2(s string, t string) bool {
	hashmap := [26]int{}
	hashmap2 := [26]int{}
	for _, v := range s {
		hashmap[v-'a']++
	}
	for _, v := range t {
		hashmap2[v-'a']++
	}
	return hashmap == hashmap2
}

// 进阶:
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
////////////////////////////////////////////
func isAnagram3(s string, t string) bool {
	hashmap := map[rune]int{}
	for _, v := range s {
		hashmap[v]++
	}
	for _, v := range t {
		hashmap[v]--
		if hashmap[v] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "rat"
	t := "card中国"
	fmt.Println(isAnagram2(s, t))
}
