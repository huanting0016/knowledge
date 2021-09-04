package main

import (
	"fmt"
)

func main() {
	var nums []int
	var res int
	fmt.Scan(&nums)
	fmt.Scan(&res)
	//fmt.Printf("%d\n", a)
	fmt.Println(search(nums, res))

}
func search(nums []int, res int) int {
	fmt.Println(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		fmt.Println(mid)
		if nums[mid] == res && nums[mid-1] != res {
			return mid
		} else if nums[mid] > res {
			right = mid - 1
		} else if nums[mid] < res {
			left = mid + 1
		} else {
			right = mid + 1
		}

	}
	return -1
}
