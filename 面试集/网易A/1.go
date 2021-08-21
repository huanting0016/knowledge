package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func target1(res []int, M int) int {
	count := 0
	for i := 0; i < len(res)-1; i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i]+res[j] <= M {
				count++
			}
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nums []int
	target := 0
	for i := 0; i < 2; i++ {
		scanner.Scan()
		if i == 0 {
			input := scanner.Text()
			strs := strings.Split(input, " ")
			nums = make([]int, len(strs))
			for i, str := range strs {
				t, _ := strconv.Atoi(str)
				nums[i] = t
			}

		} else {
			input := scanner.Text()
			target, _ = strconv.Atoi(input)
		}
	}

	fmt.Println(target1(nums, target))
}
