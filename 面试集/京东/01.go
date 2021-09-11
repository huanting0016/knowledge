package main

import "fmt"

func main() {
	var n, m, x, y, z int
	fmt.Scan(&n, &m, &x, &y, &z)
	var res [][]string
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		var res1 []string
		for j := 0; j < m; j++ {
			res1 = append(res1, string(str[j]))
		}
		res = append(res, res1)
	}
	var target string
	fmt.Scan(&target)

	fmt.Println(minTime(res, target, x, y, z))
}

func minTime(res [][]string, target string, x, y, z int) int {
	var losttime int
	var lastJ, lastK int
	for i := 0; i < len(target); i++ {

		for j := 0; j < len(res); j++ {
			for k := 0; k < len(res[0]); k++ {
				if res[j][k] == string(target[i]) {
					if j == lastJ && k == lastK {
						losttime += z
					} else if j == lastJ {
						losttime += equ(k-lastK)*x + z
					} else if k == lastK {
						losttime += equ(j-lastJ)*x + z
					} else {
						losttime += (equ(k-lastK)+equ(j-lastJ))*x + y + z
					}

					lastJ, lastK = j, k
				}
			}
		}

	}

	return losttime

}
func equ(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
