package main

import (
	"fmt"
	"strings"
)

//
//type copyline struct {
//	content string
//}
//
//var res []copyline
//var wg =sync.WaitGroup{}
//func (l copyline) fanyi(language string,i int) {
//	defer wg.Done()
//	r:=strings.ToUpper(l.content)
//	res[i] = copyline{r}
//}
//
//func fanyiAll( copylines []copyline,language string){
//	res = make([]copyline,len(copylines))
//
//	for i := 0; i < len(copylines); i++ {
//		wg.Add(1)
//		go copylines[i].fanyi(language,i)
//	}
//	wg.Wait()
//}
//
//func main() {
//	res = []copyline{}
//	str:=""
//	fmt.Scanln(&str)
//	strs:=strings.Split(str,",")
//	copylines:=make([]copyline,len(strs))
//	for i, s := range strs {
//		copylines[i] = copyline{s}
//	}
//	fanyiAll(copylines,"UPPER")
//	for i := 0; i < len(res); i++ {
//		if i!=len(res)-1{
//			fmt.Print(res[i].content,",")
//		}else {
//			fmt.Print(res[i].content)
//		}
//	}
//}

func main() {
	str := ""
	fmt.Scanln(&str)
	strs := strings.Split(str, ",")
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 || i != 0 {
			fmt.Print(",")
		} else {
			fmt.Print(strings.ToUpper(strs[i]))
		}
	}
}
