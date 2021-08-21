package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nums []int
	scanner.Scan()
	input := scanner.Text()
	strs := strings.Split(input, " ")
	nums = make([]int, len(strs))
	for i, str := range strs {
		t, _ := strconv.Atoi(str)
		nums[i] = t
	}
	if len(nums) == 1 {
		fmt.Println(1)
	} else if len(nums) == 2 {
		if nums[0] != nums[1] {
			fmt.Println(3)
		} else {
			fmt.Println(2)
		}

	} else {
		fmt.Println(5)
	}

}
