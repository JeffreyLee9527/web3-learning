package main

import (
	"fmt"
	"sync"
)

/**
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/

func main() {
	const num = 10
	var mu sync.Mutex
	var wg sync.WaitGroup
	var cnt int
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			fmt.Printf("协程:%d\n", i)
			for j := 0; j < 1000; j++ {
				cnt++
			}
		}()
	}
	wg.Wait()
	fmt.Printf("计算结束,num : %d\n", cnt)
	fmt.Printf("程序结束")
}
