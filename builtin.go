package main

import "fmt"

//内置函数就是不用导入包直接用
func main() {
	//len函数
	var num1 string
	fmt.Scanln(&num1)
	fmt.Println(len(num1))
	//new函数，分配内存，返回值是对应类型的指针，参数是具体类型
	new1 := new(int)
	fmt.Printf("new1的类型: %T,new1的值: %v,new1的地址是:%v,new1指向的值: %v", new1, new1, &new1, *new1)
	//new1的值就是new1指向的值的地址
}
