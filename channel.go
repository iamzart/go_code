package main

import (
	"fmt"
	"sync"
)

/*func main() {
	cha := make(chan string, 1)
	cha <- "hello world"
	fmt.Println(<-cha)
	defer close(cha)
}
*/

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i) //这里是匿名函数传参的结构，// 正确：每个 goroutine 拿到的都是传递过来的 i 值
	}
	wg.Wait()
}
