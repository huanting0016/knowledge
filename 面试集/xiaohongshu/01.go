package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	var nums [][]string
	var index_x, index_y int
	for i := 0; i < r; i++ {
		var num []string
		var str string
		fmt.Scan(&str)
		for j := 0; j < c; j++ {
			if str[j] == 'R' {
				index_x = i
				index_y = j
			}
			num = append(num, string(str[j]))
		}
		nums = append(nums, num)
	}
	var zhiling int
	var zhilingji []string
	fmt.Scan(&zhiling)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		zhilingji = append(zhilingji, input)
	}
	res1, res2 := pianyi(nums, index_x, index_y, zhilingji)
	fmt.Println(res1, res2)

}
func pianyi(nums [][]string, index_x, index_y int, zhilingji []string) (int, int) {
	direction := "up"
	res_x, res_y := index_x, index_y
	for i := 0; i < len(zhilingji); i++ {
		if zhilingji[i] == "Turn right" {
			switch direction {
			case "up":
				direction = "right"
			case "right":
				direction = "down"
			case "down":
				direction = "left"
			case "left":
				direction = "up"
			}
		} else if zhilingji[i] == "Turn left" {
			switch direction {
			case "up":
				direction = "left"
			case "right":
				direction = "up"
			case "down":
				direction = "right"
			case "left":
				direction = "down"
			}
		}
		if strings.Contains(zhilingji[i], "Forward") {
			zhiling_nums := strings.Split(zhilingji[i], " ")
			zhiling, _ := strconv.Atoi(zhiling_nums[1])
			if direction == "up" {
				for j := 0; j < zhiling; j++ {
					if index_x == 0 {
						break
					}
					if nums[index_x][index_y] == "O" {
						index_x += 1
						break
					}
					index_x -= 1
				}
			} else if direction == "down" {
				for j := 0; j < zhiling; j++ {
					if index_x == len(nums)-1 {
						break
					}
					if nums[index_x][index_y] == "O" {
						index_x -= 1
						break
					}
					index_x += 1
				}
			} else if direction == "right" {
				for j := 0; j < zhiling; j++ {
					if index_y == len(nums[0])-1 {
						break
					}
					if nums[index_x][index_y] == "O" {
						index_y -= 1
						break
					}
					index_y += 1
				}
			} else if direction == "left" {
				for j := 0; j < zhiling; j++ {
					if index_y == 0 {
						break
					}
					if nums[index_x][index_y] == "O" {
						index_y += 1
						break
					}
					index_y -= 1
				}
			}
		}

	}
	return index_x - res_x, index_y - res_y

}
