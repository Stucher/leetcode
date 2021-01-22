package main

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	fmt.Println(n)
	result := countAndSay(n - 1)
	fmt.Printf("result->%v\n",result)

	tmpArr := strings.Split(result,"")
	tmpRet := make([]string,0)
	c:=""
	cnt:=0

	for i := 0; i < len(tmpArr); i++ {
		if tmpArr[i]!=c && c!=""{
			tmpRet = append(tmpRet,fmt.Sprintf("%d",cnt))
			tmpRet = append(tmpRet,fmt.Sprintf("%v",c))
			cnt = 0
		}
		c = tmpArr[i]
		cnt++
	}
	tmpRet = append(tmpRet,fmt.Sprintf("%d",cnt))
	tmpRet = append(tmpRet,fmt.Sprintf("%v",c))

	fmt.Printf("%v\n",tmpRet)
	result = strings.Join(tmpRet,"")

	return result
}
