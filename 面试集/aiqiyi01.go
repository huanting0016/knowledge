package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputR := bufio.NewReader(os.Stdin)
	input, _ := inputR.ReadString('\n')
	numstring := strings.Split(input, ":")
	k, _ := strconv.Atoi(strings.TrimSpace(numstring[1]))
	nums := strings.Split(numstring[0], ",")
	var p []float64
	var maxP float64
	var sum int
	var count int

	for i := 0; i < len(nums)-k+1; i++ {

		for j := 0; j < k; j++ {

			value, _ := strconv.Atoi(strings.TrimSpace(nums[i+j]))
			sum += value
			count++
			if count == k {
				var p1 float64
				p1 = float64(sum) / float64(k)
				sum = 0
				count = 0
				p = append(p, p1)
			}

		}

	}

	for z := 0; z < len(p)-2; z++ {
		ca := (p[z+1] - p[z]) * 100 / p[z]
		maxP = max(maxP, ca)
	}
	str := strconv.FormatFloat(maxP, 'f', 2, 64)

	fmt.Println(str + "%")

}
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
