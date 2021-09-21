package main

import "fmt"

func main() {
	var n,m int
	fmt.Scan(&n,&m)
	var res []string
	for i:=0;i<m;i++{
		var q,w string
		fmt.Scan(&q,&w)
		if q > w{
			q,w = w,q
		}
		if len(res) == 0{
			res = append(res,q+w)
		}else{
			flag := false
			for j:=0;j<len(res);j++{
				if string(res[j][len(res[j])-1]) == q{
					res[j]+=w
					flag = true
				}
			}
			if !flag{
				res = append(res,q+w)
			}
		}


	}

	var count int
	for i:=0;i<len(res);i++{
		if len(res[i]) >5{
			flag1 := false
			for _,v:=range res{
				if v == string(res[i][0]) + string(res[i][len(res[i])-1]) || v ==  string(res[i][len(res[i])-1]) + string(res[i][0])  {
					count += len(res[i])
					flag1 = true
				}
			}
			if !flag1{
				count += len(res[i])-4
			}

		}else if len(res[i]) == 5{
			count+= 1
		}
	}
fmt.Println(count)
}