package main

import "fmt"

func rotate(matrix [][]int) {
	// 先进行水平翻转，再进行对角线翻转
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[len(matrix)-1-i] = matrix[len(matrix)-1-i], matrix[i]
	}
	for i := 0; i < len(matrix); i++ {
		// 矩阵从1到n，即为从0到n-1
		for j := i + 1; j < len(matrix); j++ {
			// 从a12一直到a1n
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}

	}

}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
