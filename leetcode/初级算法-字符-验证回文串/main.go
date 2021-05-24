package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	hashmap := []uint8{}
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			//fmt.Println("字母：",string(rune(s[i])))
			if s[i] >= 65 && s[i] <= 90 {
				//fmt.Println("大写")
				hashmap = append(hashmap, s[i]+uint8(32))
			} else {
				hashmap = append(hashmap, s[i])
			}
		}
	}
	fmt.Println(hashmap)
	if len(hashmap) == 0 {
		return false
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

func main() {
	s := ".,"
	fmt.Println(isPalindrome(s))
}
