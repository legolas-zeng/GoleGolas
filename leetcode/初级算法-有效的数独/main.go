package main

import "fmt"

/////////////////常规解法，分为三部分判断////////////////////////////
//func do(arr [][]string) bool {
//	for i := 0; i < 9; i++ {
//		hashmaprow := map[string]int{}
//		hashmapcol := map[string]int{}
//		hashmapshudu := map[string]int{}
//		for j := 0; j < 9; j++ {
//			//行 判断
//			if arr[i][j] != '.' {
//				if n, ok := hashmaprow[arr[i][j]]; ok {
//					fmt.Println(n, '行存在相同元素')
//					return false
//				}
//				hashmaprow[arr[i][j]] = 1
//			}
//
//			//列 判断
//			if arr[j][i] != '.' {
//				if n, ok := hashmapcol[arr[j][i]]; ok {
//					fmt.Println(n, '列存在相同元素')
//					return false
//				}
//				hashmapcol[arr[j][i]] = 1
//
//			}
//			//数独判断
//			row := (i%3)*3 + j%3
//			col := (i/3)*3 + j/3
//			if arr[row][col] != '.' {
//				if _, ok := hashmapshudu[arr[row][col]]; ok {
//					fmt.Println(row, col, '数独存在相同元素')
//					return false
//				}
//				hashmapshudu[arr[row][col]] = 1
//
//			}
//		}
//	}
//	return true
//}

func main() {
	//arr := [][]string{
	//	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	//	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	//	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	//	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	//	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	//	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	//	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	//	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	//	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	//}
	arr1 := [][]byte{
		{'5', '3', '.', '1', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(do(arr1))
	//fmt.Println(1 % 3)
}

func do(arr [][]byte) bool {
	var row, col [9][9]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i][j] != '.' {
				num := arr[i][j] - '1'
				if row[i][num] == 1 {
					// 如果出现个重复的数字，他们row的索引肯定会有记录，第二次读取发现等于1，原来有记录，就返回flase
					return false
				} else {
					// 只要行没有出现这个arr[i][j]数字，那个arr[i][j]-1也就是唯一的，把它当做row的行索引也是唯一的。
					row[i][num] = 1
				}
				if col[j][num] == 1 {
					return false
				} else {
					col[j][num] = 1
				}
			}

		}
	}

	return true
}
