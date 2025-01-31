package main

import "fmt"

func Getsum() func(int) int {
	x := 10
	return func(n int) int {
		x += n
		return x
	}

}
func main() {
	f1 := Getsum()     //f1是一个匿名函数，因为getsum的返回值是匿名函数，而f1来接收的
	fmt.Println(f1(1)) //f1(1)相当于给这个匿名函数传的参数是1，

}
