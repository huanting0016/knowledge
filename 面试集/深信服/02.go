package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param grid int整型二维数组
 * @return int整型
 */
type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func ncov_defect(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}
	// write code here
	var count int
	n, m := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				for _, d := range dir4 {
					if x, y := i+d.x, j+d.y; x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 0 {
						count++
					}
				}
			}
		}
	}
	return count
}
