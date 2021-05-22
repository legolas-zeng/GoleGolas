package main

import "fmt"

//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

//输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

func rotate(matrix [][]int) {
	lenarr := len(matrix[0])
	for i := 0; i < lenarr; i++ {
		tmp := []int{}
		for j := 0; j < lenarr; j++ {
			tmp = append(tmp, matrix[lenarr-j-1][i])
		}
		fmt.Println(tmp)
	}
}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
}
