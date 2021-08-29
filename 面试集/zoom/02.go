package main

import (
	"fmt"
	"math"
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @param windowSize int整型 滑动窗口大小
 * @param step int整型 每次移动step个数字
 * @return int整型一维数组
 */
func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(slideWindow(nums, 11, 2))
}

func slideWindow(nums []int, windowSize int, step int) []int {
	// write code here
	var res []int
	if len(nums) <= windowSize {
		sort.Ints(nums)
		res = append(res, nums[len(nums)-1])
		return res
	}
	for i := 0; i < len(nums); i += step {
		maxNum := math.MinInt64
		for j := i; j < i+windowSize && j < len(nums); j++ {
			if nums[j] > maxNum {
				maxNum = nums[j]
			}
		}
		res = append(res, maxNum)
	}
	return res
}
