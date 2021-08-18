package main

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶(n为正整数)总共有多少种跳法

输入：3
返回值：4
*/

func jumpFloorII(number int) int {
	if number <= 2 {
		return number
	}
	f0 := 1
	f1 := 1
	f2 := 2
	sum := f0 + f1 + f2
	if number == 3 {
		return sum
	} else {
		for i := 4; i <= number; i++ {
			sum *= 2
		}
	}

	return sum
}
