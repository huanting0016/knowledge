package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var j, k int
		fmt.Scan(&j, &k)
		scanner := bufio.NewScanner(os.Stdin)
		var nums []int
		scanner.Scan()
		input := scanner.Text()
		strs := strings.Split(input, " ")
		nums = make([]int, len(strs))
		for q, str := range strs {
			t, _ := strconv.Atoi(str)
			nums[q] = t
		}

		fmt.Println(people(nums, k))
	}

}
func people(nums []int, index int) int {
	var count1, count2 int
	flag1, flag2 := true, true
	index_min1, index_min2 := math.MaxInt64, math.MaxInt64
	var index_max1, index_max2 int
	fmt.Println(flag2, index_min2, index_max2)
	for i := index; i < len(nums); {

		if nums[i] > nums[index-1] && flag1 {
			index_max1 = i
			flag1 = false
			i++
			continue
		}
		if nums[i] < nums[index_max1] && !flag1 {
			index_min1 = i

		} else if nums[i] > nums[index_max1] && !flag1 {
			index_max1 = i

		}
		if index_min1 < index && nums[i] > nums[index_min1] && !flag1 {

			index_max1 = i
			count1++
		}
		if i == len(nums)-1 && nums[i] < nums[index_max1] {
			count1++
		}

		i++
	}

	//for i:=index-1;i>=0;{
	//
	//
	//	if nums[i] > nums[index-1] && flag2{
	//		index_max2 = i
	//		flag2 = false
	//		i--
	//		continue
	//	}
	//	if nums[i] < nums[index_max2] && !flag2{
	//		index_min2 = i
	//
	//
	//	}else{
	//		index_max2 = i
	//
	//	}
	//	if nums[i] > nums[index_min2] && !flag2{
	//		index_max2 = i
	//		count2++
	//	}
	//	if i == 0 && nums[i] < nums[index_max2]{
	//		count2++
	//	}
	//
	//	i--
	//}
	fmt.Println(count1, count2)
	return count1 + count2
}
