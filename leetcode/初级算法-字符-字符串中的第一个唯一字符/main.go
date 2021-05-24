package main

import "fmt"

//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//
//示例：
//s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2
///////////////////////使用map  hash表////////////////////////////////
func firstUniqChar(s string) int {
	hashmap := map[string]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := hashmap[string(rune(s[i]))]; ok {
			fmt.Println("有重复：", string(rune(s[i])))
			hashmap[string(rune(s[i]))] = 1
		} else {
			fmt.Println("存入", string(rune(s[i])), i)
			hashmap[string(rune(s[i]))] = 0
		}
	}
	fmt.Println(hashmap)
	//for i := 0; i < len(s); i++ {
	//	if hashmap[string(rune(s[i]))] == 0 {
	//		fmt.Println("第一个不重复", string(rune(s[i])))
	//		return i
	//	}
	//}
	for k, v := range s {
		if hashmap[string(rune(v))] == 0 {
			return k
		}
	}
	return -1
}

///////////////////////使用切片 hash表/////////////////////////////////
func firstUniqChar2(s string) int {
	hash := [26]int{}
	for i := range s {
		hash[s[i]-'a']++
	}
	fmt.Println(hash)
	for k, v := range s {
		if hash[v-'a'] == 1 {
			return k
		}
	}
	return -1
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}
