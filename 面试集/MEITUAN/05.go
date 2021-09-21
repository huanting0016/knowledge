package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
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


}
