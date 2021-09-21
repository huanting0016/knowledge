package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var k int
	fmt.Scan(&k)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strs := strings.Split(input, " ")
	var nums []int
	var nums1 []int
	for i := 0; i < len(strs); i++ {
		t, _ := strconv.Atoi(strs[i])
		nums = append(nums, t)
		nums1 = append(nums, t)
	}
	sort.Ints(nums)
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums1[i] {
			continue
		}
		count++
	}
	if count%2 == 0 {
		fmt.Println(count / 2)
	} else {
		fmt.Println(count/2 + 1)
	}
}
