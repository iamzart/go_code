package main

import "fmt"

//key不能重复
/*func main() {
//map要用make来分配内存，只声明没用
nums := make(map[int]string, 5)
// for i := 0; i < 5; i++ {
// 	fmt.Scanf("%d", &key)
// 	fmt.Scanf("%s", &name)
// 	//一一对应
// 	nums[key] = name
// }

//将键值对存入map中
nums[1] = "张三"
nums[2] = "李四"
nums[3] = "王五"
nums[4] = "赵六"
nums[5] = "田七"
for key, name := range nums {
	fmt.Println(key, name)
}
fmt.Println(nums)
fmt.Println(nums[2])
*/
/*增删查改*/
/*
	b := map[int]string{
		1: "张山",
		2: "lisa",
		3: "rose",
		4: "jack",
		5: "sum",
	}
	//增
	b[6] = "lili"
	//删
	//delete函数，第一个参数是一个map变量，第二个是对应的键
	delete(b, 3)
	//改
	b[2] = "mike"
	//查
	fmt.Println(b[2])
	fmt.Println(b)

	//清空
	//clear(b)

	fmt.Println(len(b))

}
*/
func main() {
	//map的嵌套
	freshman := map[string]map[int]string{
		"软件工程": {
			01: "lisa",
			02: "jack",
			03: "mike",
		},
		"大数据": {
			01: "lili",
			02: "jack",
			03: "mike",
		},
		"人工智能": {
			01: "lili",
			02: "jack",
			03: "mike",
		},
	}
	fmt.Println(freshman)
}
