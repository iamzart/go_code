package main

import (
	"fmt"
	"time"
)

func Add(start, end int, ch chan int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	ch <- sum
}

func Div(a int, b int) {

	defer func() {
		e := recover()
		if e != nil {
			fmt.Println(e)
		}
	}()
	fmt.Println(a / b) //要先捕捉再处理
}

func main() {
	ch := make(chan int, 4)
	go Add(1, 50, ch)
	go Add(51, 100, ch)
	//此时ch里存了两次的数据
	total := <-ch + <-ch
	fmt.Println(total)
	go Div(1, 0)
	go Div(6, 3)

	time.Sleep(time.Second * 5) //让主线程等等协程
}
