package main

import "fmt"

func tmpAraay(nums []int) int {
	res := make(map[int]int)
	tmp := make(map[int]int)
	max1, max2 := 0, 0
	for i := 0; i < len(nums); i++ {
		// 偶数
		if i%2 == 0 {
			tmp[nums[i]] += 1
			max2 = max(max2, tmp[nums[i]])
		}
		if i%2 == 1 {
			res[nums[i]] += 1
			max1 = max(max1, res[nums[i]])
		}
	}
	return len(nums) - max1 - max2
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	res := tmpAraay(nums)
	fmt.Println(res)
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
