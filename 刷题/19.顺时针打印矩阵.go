package main

/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，例如，如果输入如下4 X 4矩阵

[[1,2,3,4],
[5,6,7,8],
[9,10,11,12],
[13,14,15,16]]

则依次打印出数字
[1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	var res []int
	l, t, r, b := 0, 0, len(matrix[0])-1, len(matrix)-1

	for {
		// 左-->右打印，打印完t++
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		if t > b {
			break
		}
		// 上-->下打印，打印完r--
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		if l > r {
			break
		}
		// 右-->左打印，打印完b--
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b--
		if t > b {
			break
		}
		// 下-->上打印
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}
	return res
}
