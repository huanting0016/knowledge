package main

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix)-1, len(matrix[0])-1
	if m <= 1 && n <= 1 {
		if matrix[0][0] == '0' {
			return 0
		} else {
			return 1
		}
	}

	count := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' && matrix[i+1][j] == '1' && matrix[i][j+1] == '1' && matrix[i+1][j+1] == 1 {
				count++
			} else {
				continue
			}

		}

	}
	for k := 0; k <= n; k++ {
		if matrix[][] == '1' && count<=1{
			count = 1
			}else {
			count = 0
			}
}
		return count * count
}
