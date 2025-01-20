package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//两个操作函数
func WithName(name string) func(*Person) {
	return func(p *Person) { //括号里是参数，变量类型写后边，前边是结构体变量
		p.Name = name
	}
}
func WithAge(age int) func(*Person) { //with函数参数是age，返回的是匿名函数，匿名函数的参数是结构体类型
	return func(p *Person) {
		p.Age = age
	}
}

func main() {
	per1 := &Person{} //per1是结构体指针类型
	option := []func(*Person){
		WithName("Alice"),
		WithAge(30),
	}
	for _, opt := range option { //opt是option（切片）中的每个元素，也就是每个with函数
		opt(per1)
	}
	fmt.Printf("name: %s,Age: %d\n", per1.Name, per1.Age)
}
