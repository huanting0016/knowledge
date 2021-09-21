package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)
	var res []string
	for i:=0;i<m;i++{
		var res1 string
		fmt.Scan(&res1)

res = append(res,res1)

	}

	for i:=0;i<len(res);i++{
		flag := false
		nums := make(map[string]bool,len(res))
		nums[string(res[i][0])] = true
		for j:=1;j<len(res[i]);j++{
			if _,ok := nums[string(res[i][j])];!ok{
				flag = true
				break
			}else{
				flag = false
				nums[string(res[i][j])] = true
			}
		}
		if flag{
			fmt.Println("YES")
		}else {
			fmt.Println("NO")
		}

	}


}
