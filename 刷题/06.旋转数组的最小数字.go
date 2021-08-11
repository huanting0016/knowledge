package main

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0

输入：[3,4,5,1,2]
返回值：1
*/

func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	if len(rotateArray) < 1 {
		return 0
	}
	if len(rotateArray) == 1 {
		return rotateArray[0]
	}
	l, r := 0, len(rotateArray)-1

	for l < r {
		if rotateArray[l] < rotateArray[r] {
			return rotateArray[l]
		}
		mid := l + (r-l)>>1
		if rotateArray[mid] > rotateArray[r] {
			l = mid + 1
		} else if rotateArray[mid] < rotateArray[r] {
			r = mid
		} else if rotateArray[mid] == rotateArray[r] {
			r = r - 1
		}
	}
	return rotateArray[l]
}
