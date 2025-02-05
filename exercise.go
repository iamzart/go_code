package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fetchURL(ctx context.Context, url string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()                                      //减少计数
	delay := time.Duration(rand.Intn(5)+1) * time.Second //随机生成1-5秒的延迟
	select {
	case <-time.After(delay):
		ch <- fmt.Sprintf("爬取成功： %s(耗时%v)", url, delay)
	case <-ctx.Done():
		ch <- fmt.Sprintf("爬取超时: %s", url)
	}
}

func main() {
	rand.Seed(time.Now().UnixMicro()) //确保随机数不同
	urls := []string{"https://example1.com", "https://example2.com", "https://example3.com"}
	ch := make(chan string, len(urls))
	var wg sync.WaitGroup
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(ctx, url, ch, &wg)
	}

	// 等待所有爬取任务完成
	go func() {
		wg.Wait()
		close(ch) // 关闭 channel，通知 main 协程不再有数据写入
	}()

	// 读取爬取结果
	for result := range ch {
		fmt.Println(result)
	}

	/*range ch 的本质是：一直等着 ch 里有数据可以读。
	如果 ch 里没有数据，它就会卡住（阻塞），等着新数据到来。
	只有 ch 关闭（close(ch)）了，range 才会自动退出。*/

	fmt.Println("🚀 所有任务处理完毕！")
}

/*搞一个爬取函数，传入参数分别是，一个ctx，就是判断超时机制，然后网址，然后管道，把爬出来的数据传进去，还有wg，控制别提前结束*/
