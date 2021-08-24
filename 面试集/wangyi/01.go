package  main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	var m int
	fmt.Scan(&m)
	for i:=0;i<m;i++{
		var inputstring string
		fmt.Scan(&inputstring)
		if len(inputstring) != 18{
			fmt.Println(0)
		}else{
			num := haveresult(inputstring)
			fmt.Println(num)
		}

	}
}
var nums = []int{7,9,10,5,8,4,2,1,6,3,7,9,10,5,8,4,2}
var index = map[string]int{"1":0,"0":1,"X":2,"9":3,"8":4,"7":5,"6":6,"5":7,"4":8,"3":9,"2":10 }
var tmd = map[int]int{0:1,2:1,4:2,6:2,8:4,10:4,12:6,14:6,16:9,18:9,20:11,22:11,24:14,26:14,28:16,30:16,32:19,34:19,36:21,38:21,40:23,42:23,44:24,46:24,48:25,50:25,52:25,54:25,56:25,58:25,60:25,62:25,64:25,66:25,68:25,70:25,72:25,74:25,76:25,78:25,80:24,82:24,84:23,86:23,88:21,90:21,92:19,94:19,96:16,98:16,100:14,102:14,104:11,106:11,108:9,110:9,112:6,114:6,116:4,118:4,120:2,122:2,124:1,126:1}
func haveresult(res string)int{
	res = strings.ToUpper(res)
	var sum,yushu,count int
	var m []int
	yushu = index[string(res[len(res)-1])]

	for i:=0;i<len(res)-1;i++{
		if (res[i] > '9' || res[i] < '0') && res[i] !='*'{
			return 0
		}
		if res[i] == '*'{
			if i>=14 && i<=16{
				m = append(m,nums[i])
			}else{
				return 0
			}

		}else{
			num,_ := 	strconv.Atoi(string(res[i]))
			sum  += nums[i]*num
		}
	}
	if len(m) == 1{
		return 1
	}else if len(m) == 2{
		for q:=0;q<=9;q++{
			for w:=0;w<=9;w++{
				if (sum + m[0] * q+m[1] * w) % 11 == yushu{
					count++
				}
			}}}else if len(m) == 3{

		for k,v :=range tmd{
			if (sum + k) % 11 == yushu{
				count += v
			}


		}

	}else if len(m)==0{
		if sum  % 11 == yushu{
			count++
		}
	}else{
		return 0
	}

	return count

}