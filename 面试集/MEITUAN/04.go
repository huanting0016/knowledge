//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//
//func main() {
//	c:=0
//	fmt.Scan(&c)
//	res:=make([]string,0,c)
//	for i := 0; i < c; i++ {
//		l:=0
//		fmt.Scan(&l)
//		nums:=""
//		fmt.Scan(&nums)
//		res = append(res, sndmeesage(nums))
//	}
//	for _, re := range res {
//		fmt.Println(re)
//	}
//}
//
//
//func sndmeesage(nums string)(res string){
//
//	ss:=strings.Split(nums,"--")
//	for _, s := range ss {
//		if len(s) == 0{
//			continue
//		}
//		s = strings.TrimLeft(s,"-")
//		s = strings.Join(strings.Split(s,"-"),"")
//		l:=len(s)
//		for i := 0; i < l;  {
//			p:=s[i]
//			c,j:=1,i+1
//			if j==l{
//				res+= getMessage(p,c)
//				break
//			}
//			for ; j <l ; j++ {
//
//				if s[j] == p{
//					c++
//					if j==l-1{
//						i = j+1
//						res+= getMessage(p,c)
//						break
//					}
//				}else{
//					i = j
//					res+= getMessage(p,c)
//					break
//				}
//			}
//
//		}
//	}
//	return
//
//}
//
//func getMessage(num byte,count int) (msg string) {
//	if num == '0'{
//		for i := 0; i < count; i++ {
//			msg+="-"
//		}
//	}else{
//		ss:= res[num]
//		s:=(count-1)%len(ss)
//		msg= ss[s]
//	}
//
//	return
//}
//
//
//var res = map[byte][]string{
//	'2':{"a","b","c"},
//	'3':{"d","e","f"},
//	'4':{"g","h","i"},
//	'5':{"j","k","l"},
//	'6':{"m","n","o"},
//	'7':{"p","q","r","s"},
//	'8':{"t","u","v"},
//	'9':{"w","x","y","z"},
//}