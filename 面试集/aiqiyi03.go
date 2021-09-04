package main

import (
	"fmt"
	"strconv"
	"strings"
)

func flood(rains []int) []int {
	sums := make([]int, len(rains))
	n := make([][]int, 0)
	m := make(map[int]int, 0)
	isFlood := false
	j := -1
	num := make([]int, 0)
	for i := 0; i < len(rains); i++ {
		if rains[i] > 0 {
			isFlood = true
		} else {
			if isFlood {
				if j > -1 {
					n = append(n, num)
				}
				num = make([]int, 0)
				j++
				isFlood = false
			}
			if j > -1 {
				num = append(num, i)
			}
		}
	}
	if len(num) > 0 {
		n = append(n, num)
	}
	j = 0
	isFlood = false
	for i := 0; i < len(rains); i++ {
		if rains[i] > 0 {
			isFlood = true
			sums[i] = -1
			if v, ok := m[rains[i]]; ok {
				if v == j {
					return []int{}
				}
				t := true
				for k := v; k < len(n); k++ {
					if len(n[k]) > 0 {
						sums[n[k][0]] = rains[i]
						n[k] = n[k][1:]
						t = false
						break
					}
				}
				if t {
					return []int{}
				}
			}
			m[rains[i]] = j
		} else {
			if isFlood {
				j++
				isFlood = false
			}
			sums[i] = 1
		}
	}
	return sums
}

func main() {
	str := ""
	fmt.Scanln(&str)
	strs := strings.Split(str[1:len(str)-1], ",")
	nums := make([]int, len(strs))
	for i, s := range strs {
		nums[i], _ = strconv.Atoi(s)
	}
	res := flood(nums)
	for i, re := range res {
		if i == 0 {
			fmt.Print("[", re, ",")
		} else if i == len(res)-1 {
			fmt.Print(re, "]")
		} else {
			fmt.Print(re, ",")
		}
	}
}
