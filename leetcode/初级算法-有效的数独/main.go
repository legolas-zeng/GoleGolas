package main

import "fmt"

func do(arr [][]string) bool {
	for i := 0; i < 9; i++ {
		hashmaprow := map[string]int{}
		hashmapcol := map[string]int{}
		hashmapshudu := map[string]int{}
		for j := 0; j < 9; j++ {
			//行 判断
			if arr[i][j] != "." {
				if n, ok := hashmaprow[arr[i][j]]; ok {
					fmt.Println(n, "行存在相同元素")
					return false
				}
				hashmaprow[arr[i][j]] = 1
			}

			//列 判断
			if arr[j][i] != "." {
				if n, ok := hashmapcol[arr[j][i]]; ok {
					fmt.Println(n, "列存在相同元素")
					return false
				}
				hashmapcol[arr[j][i]] = 1

			}
			//数独判断
			row := (i%3)*3 + j%3
			col := (i/3)*3 + j/3
			if arr[row][col] != "." {
				if _, ok := hashmapshudu[arr[row][col]]; ok {
					fmt.Println(row, col, "数独存在相同元素")
					return false
				}
				hashmapshudu[arr[row][col]] = 1

			}
		}
	}
	return true
}

func main() {
	arr := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	fmt.Println(do(arr))
	//fmt.Println(1 % 3)
}
