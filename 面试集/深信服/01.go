package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param a string字符串 字符串1
 * @param b string字符串 字符串2
 * @return int整型
 */
func minDistance(a string, b string) int {
	// write code here
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j

	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minint(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
func minint(args ...int) int {
	min := args[0]
	for _, item := range args {
		if item < min {
			min = item
		}
	}
	return min

}
