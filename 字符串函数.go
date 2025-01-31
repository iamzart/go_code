package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "我爱你中国"

	//求字符串的长度 len，一个汉字三个字节
	fmt.Println(len(str))

	//对字符串进行遍历
	//方式一：用for range
	for i, val := range str {
		fmt.Printf("索引为:%d,值为:%c \n", i, val)
	} //用格式化的形式用printf
	//方式二：转成切片
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}
	fmt.Println()

	//字符串转整数
	num1, err := strconv.Atoi(str)
	num2, err := strconv.Atoi("6669988")

	if err != nil {
		fmt.Println("转换错误", err)
	}
	fmt.Println(num1)
	fmt.Println(num2)

	//整数转字符串
	var num3 int
	fmt.Scanln(&num3)
	str1 := strconv.Itoa(num3)
	fmt.Println(str1)

	//统计字符串中有几个子串
	count := strings.Count("abaaabbcabcabbcbacb", "ab")
	fmt.Println(count)

	//比较两个字符串
	/*不区分大小写*/
	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)
	/*区分大小写*/
	fmt.Println("hello" == "HELLO")

	//返回子串在字符串中第一次出现的索引数，没有返回-1
	id := strings.Index("abcdabafbhdc", "ab")
	fmt.Println(id)

}
