package main

/*
给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。

保证base和exponent不同时为0。不得使用库函数，同时不需要考虑大数问题，也不用考虑小数点后面0的位数

输入：2.00000,3
返回值：8.00000

输入：2.00000,-2
返回值：0.25000
说明：2的-2次方等于1/4=0.25
*/

func Power(base float64, exponent int) float64 {
	// write code here
	var res = 1.0
	if exponent == 0 {
		return res
	}
	if exponent < 0 {
		base = 1.0 / base
		exponent = -exponent
	}

	for exponent != 0 {
		if exponent%2 == 1 {
			res *= base
		}
		base = base * base
		exponent = exponent / 2
	}

	return res
}
