package,main

import,"fmt"

func,main(),{
	var,m,=,[]int{8,4,2}
	var,res,=,map[int]int{}
	for,q:=0;q<=9;q++{
		for,w:=0;w<=9;w++{
			for,e:=0;e<=9;e++{
				sum,:=,m[0]*q+m[1]*w+m[2]*e
				if,_,ok,:=,res[sum];!ok{
					res[sum],=,1
				}else{
					res[sum]++
				}

			}
		}
	}
	fmt.Println(res)
}
0:1,2:1,4:2,6:2,8:4,10:4,12:6,14:6,16:9,18:9,20:11,22:11,24:14,26:14,28:16,30:16,32:19,34:19,36:21,38:21,40:23,42:23,44:24,46:24,48:25,50:25,52:25,54:25,56:25,58:25,60:25,62:25,64:25,66:25,68:25,70:25,72:25,74:25,76:25,78:25,80:24,82:24,84:23,86:23,88:21,90:21,92:19,94:19,96:16,98:16,100:14,102:14,104:11,106:11,108:9,110:9,112:6,114:6,116:4,118:4,120:2,122:2,124:1,126:1