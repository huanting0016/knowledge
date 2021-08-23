package main

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，
并保证奇数和奇数，偶数和偶数之间的相对位置不变。

输入：[1,2,3,4]
返回值：[1,3,2,4]

输入：[2,4,6,5,7]
返回值：[5,7,2,4,6]
*/

func reOrderArray(array []int) []int {
	var res1, res2 []int
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			res1 = append(res1, array[i])
		} else {
			res2 = append(res2, array[i])
		}
	}
	res2 = append(res2, res1...)
	return res2
}
