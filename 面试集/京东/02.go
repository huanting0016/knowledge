package main

import "fmt"

var (
	req   = make(map[int]map[int]bool)
	red   = make(map[int]map[int]bool)
	opens = make(map[int]bool)
)

func open(id int) {
	if opens[id] {
		return
	}
	for de, _ := range req[id] {
		open(de)
	}
	opens[id] = true
}

func close(id int) {
	if !opens[id] {
		return
	}
	for rd, _ := range red[id] {
		close(rd)
	}
	delete(opens, id)

}

func inits(id int, dep []int) {
	if _, ok := req[id]; !ok {
		req[id] = map[int]bool{}
	}
	for _, i := range dep {
		req[id][i] = true
		if _, ok := red[i]; !ok {
			red[i] = map[int]bool{}
		}
		red[i][id] = true
	}
}

func opr(x, y int) int {
	switch x {
	case 0:
		close(y)
	case 1:
		open(y)
	}
	return len(opens)
}

func main() {
	req = make(map[int]map[int]bool)
	red = make(map[int]map[int]bool)
	opens = make(map[int]bool)
	n, q := 0, 0
	fmt.Scan(&n, &q)
	for i := 0; i < n; i++ {
		c := 0
		fmt.Scan(&c)
		des := make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&des[j])
		}
		inits(i+1, des)
	}
	res := make([]int, 0, q)
	for i := 0; i < q; i++ {
		x, y := 0, 0
		fmt.Scan(&x, &y)
		res = append(res, opr(x, y))
	}
	for _, re := range res {
		fmt.Println(re)
	}
}
