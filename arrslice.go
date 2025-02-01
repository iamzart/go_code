package main

import "fmt"

//数组的长度属于类型的一部分
// func main() {
// 	var nums [5]int
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("请录入第%d个学生的成绩： ", i+1)
// 		fmt.Scanln(&nums[i])
// 		sum += nums[i]
// 	}
// 	ave := float64(sum) / 5.0
// 	fmt.Println("sum = ", sum)
// 	fmt.Println("ave = ", ave)
// 	fmt.Println(nums)
// 	//获取指针
// 	pnums := new([5]int)
// 	p2nums := &nums
// 	fmt.Println(*pnums)
// 	fmt.Printf("%p\n", p2nums)
// 	fmt.Println(len(nums))
// 	//数组名是首元素的地址，数组内存是连续的

// 	/*for range遍历*/
// 	for i, val := range nums {
// 		fmt.Println(i+1, val)
// 	}
// 	for _, val := range nums {
// 		fmt.Println(val) //想忽略某个值用下划线代替
// 	}

// 	/*数组初始化*/
// 	//方式1
// 	//var arr1 = [3]int{1,2,3}
// 	//方式2
// 	//var arr2 = [...]int{1,2,3}
// 	test(&nums)
// 	fmt.Println(nums)
// }

// func test(arr *[5]int) {
// 	(*arr)[0] = 999

// 	/*二维数组*/
// }

/*切片*/

func main() {
	//创建方式一
	var arr = [5]int{5, 4, 3, 2, 1}
	var slice = arr[1:4] //前闭后开
	//slice := arr[1:4]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice)) //自动获取容量，可以自动扩容
	//创建方式二
	slice1 := make([]int, 4, 6)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)
	for i := 0; i < len(slice1); i++ {
		fmt.Scanln(&slice1[i])
	}
	for j := 0; j < len(slice1); j++ {
		fmt.Printf("slice[%v] = %v\n", j, slice1[j])
	}
	for id, val := range slice1 {
		fmt.Printf("slice[%v] = %v\n", id, val)
	}
}
