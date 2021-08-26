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
	numstring := strings.Fields(input)

	var res []int
	for i := 0; i < len(numstring)-1; i++ {
		a1, _ := strconv.Atoi(numstring[i])
		a2, _ := strconv.Atoi(numstring[i+1])
		if a1 > a2 {
			res = append(res, i+1)
		} else {
			continue
		}
	}
	fmt.Println(res[0], res[1]+1)

}
