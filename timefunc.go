package main

import (
	"fmt"
	"time"
)

/*时间函数：导入time包，使用Now函数*/
//Now函数返回的是一个结构体，结构体里又有好多固定的成员
func main() {
	now := time.Now() //函数是包里的函数
	fmt.Printf("年月日：%v %v %v 时分秒:%v %v %v\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//获得一个字符串，上面只是把她以字符串的形式直接打印出来，但是找不到这个字符串
	datestr := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(datestr)

	//格式化时间
	datestr2 := now.Format("01 02 15 04")
	fmt.Println(datestr2)
}
