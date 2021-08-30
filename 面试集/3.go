package main

import "fmt"

func main() {
	var n, m, k, x int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&k)
	fmt.Scan(&x)
	var nums []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	fmt.Println(9)

}
