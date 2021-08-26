package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var j, k int
		fmt.Scan(&j)
		fmt.Scan(&k)
		var xZ [4]int
		for q := 0; q < 4; q++ {
			var w int
			fmt.Scan(&w)
			xZ[q] = w
		}
		var qipan [][]string
		for t := 0; t < j; t++ {
			var qipan1 []string
			for e := 0; e < k; e++ {
				var r string
				fmt.Scan(&r)
				qipan1 = append(qipan1, r)
			}
			qipan = append(qipan, qipan1)

		}
		t := conTime(qipan)
		fmt.Println((xZ[0] * xZ[2]) ^ t ^ (xZ[1] * xZ[3]))
	}
}

func conTime(nums [][]string) int {

}
