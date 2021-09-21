package main

import (
	"bufio"
	"fmt"
	"os"
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
	for i := 0; i < len(strs); i++ {
		t, _ := strconv.Atoi(strs[i])
		nums = append(nums, t)
	}
	pre := make([]int, k)
	pre1 := predoing(nums, pre)
	f := make([]int, k)
	var j int
	for i := 3; i <= k-1; i++ {
		if f[i] >= f[i-1] {
			f[i] = f[i-1]
		}
		if j == ok(i, pre1) && j != 0 {
			if f[i] >= f[j-1]+1 {
				f[i] = f[j-1] + 1
			}
		}
	}
	fmt.Println(f[k-1])
}

func predoing(nums []int, pre []int) []int {
	for i := 0; i <= len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i] == nums[j] {
				pre[j] = i
				break
			}
		}
	}
	return pre
}

func ok(i int, pre []int) int {
	pos := pre[i]
	var posJ, lenM int
	for j := i - 1; j >= 0; j-- {
		if j == pos {
			continue
		}
		posJ = pre[j]
		if lenM >= posJ {
			lenM = posJ
		}

	}
	if lenM > pos {
		lenM = pos
	}
	return lenM
}
