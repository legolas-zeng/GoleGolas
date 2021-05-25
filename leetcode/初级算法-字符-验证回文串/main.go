package main

import (
	"fmt"
	"unicode"
)

/////////////////////使用双指针////////////////////////////
//先处理字符串，去除特殊字符，保留字母和数字，大小写转换
func isPalindrome(s string) bool {
	hashmap := []uint8{}
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			fmt.Println("字母：", string(rune(s[i])))
			if s[i] >= 65 && s[i] <= 90 {
				fmt.Println("大写")
				hashmap = append(hashmap, s[i]+uint8(32))
			} else {
				hashmap = append(hashmap, s[i])
			}
		}
	}
	fmt.Println(hashmap)
	if len(hashmap) == 0 {
		return true
	}
	left, right := 0, len(hashmap)-1
	for i := 0; i < len(s)/2; i++ {
		if hashmap[left] == hashmap[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

////////////////////使用双指针原地判断/////////////////////
func isPalindrome2(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
			left++
		} else if !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
			right--
		} else {
			if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
				return false
			}
			left++
			right--
		}
	}
	return true
}

//判断是否是字母数字
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main() {
	s := ".,"
	fmt.Println(isPalindrome2(s))
}
