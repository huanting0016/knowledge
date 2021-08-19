package main

/*
输入一个整数，输出该数32位二进制表示中1的个数。其中负数用补码表示

输入：10
返回值：2
*/

func NumberOf1(n int) int {
	// write code here
	var count int
	for n != 0 {
		n = n & int(int32(n)-1)
		count++
	}
	return count
}
