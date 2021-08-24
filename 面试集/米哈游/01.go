package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var res string
		fmt.Scan(&res)
		if isTrue(res) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
func isTrue(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var a, b int
	for i := 0; i < len(s); i++ {
		if b > a {
			return false

		}
		if s[i] == 'a' {
			a++
		} else {
			b++
		}
	}
	if a == b {
		return true
	}
	return false
}
