package main

import (
	"fmt"
	"math"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算最小航行费用
 * @param input int整型二维数组 二维网格
 * @return int整型
 */
func main() {
	input := [][]int{{1, 1, 1, 1, 0}, {0, 1, 0, 1, 0}, {1, 1, 2, 1, 1}, {0, 2, 0, 0, 1}}

	minSailCost(input)
}

func minSailCost(input [][]int) int {
	input[0][0] = 0
	for i := 1; i < len(input); i++ {
		isNot := true
		if input[i][0] == 0 && isNot {
			input[i][0] = input[i-1][0] + 2
		} else if input[i][0] == 1 && isNot {
			input[i][0] = input[i-1][0] + 1
		} else {
			input[i][0] = math.MaxInt64 / 2
		}
	}
	fmt.Println(input)
	for j := 1; j < len(input[0]); j++ {
		isNot := true
		if input[0][j] == 0 && isNot {
			input[0][j] = input[0][j-1] + 2
		} else if input[0][j] == 1 && isNot {
			input[0][j] = input[0][j-1] + 1
		} else {
			input[0][j] = math.MaxInt64 / 2
			isNot = false
		}
	}
	fmt.Println(input)

	for i := 1; i < len(input); i++ {
		for j := 1; j < len(input[0]); j++ {
			if input[i][j] == 0 {
				input[i][j] = min(input[i-1][j], input[i][j-1]) + 2
			} else if input[i][j] == 1 {
				input[i][j] = min(input[i-1][j], input[i][j-1]) + 1
			} else {
				input[i][j] = math.MaxInt64 / 2
			}

		}
	}
	fmt.Println(input)
	if input[len(input)-1][len(input[0])-1] >= math.MaxInt64/2 {
		return -1
	}
	return input[len(input)-1][len(input[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
