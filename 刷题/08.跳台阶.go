package main

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）

示例一：
输入：2
返回值：2


示例二：
输入：7
返回值：21

*/

func jumpFloor(number int) int {
	f1 := 1
	f2 := 2
	if number <= 2 {
		return number
	}
	var sum int
	for i := 3; i <= number; i++ {
		sum = f1 + f2
		f1 = f2
		f2 = sum
	}
	return sum

}
