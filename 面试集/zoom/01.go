package main

func find132Pattern(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	var count int
	// write code here
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[k] && nums[k] < nums[j] && i < j && j < k {
					count++
				}
			}
		}
	}
	return count
}
