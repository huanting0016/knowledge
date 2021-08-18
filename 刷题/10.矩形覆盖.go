package main

/*
我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，从同一个方向看总共有多少种不同的方法？

比如n=3时，2*3的矩形块有3种不同的覆盖方法(从同一个方向看)

输入：4    2*1的小矩形的总个数n
返回值：5   覆盖一个2*n的大矩形总共有多少种不同的方法(从同一个方向看)
*/

func rectCover(number int) int {
	if number <= 2 {
		return number
	}
	a, b := 1, 2
	var sum int
	for i := 3; i <= number; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}
