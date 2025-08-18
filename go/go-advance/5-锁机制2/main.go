package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func main() {
	const num = 10
	const num2 = 1000
	var cnt int64
	var wg sync.WaitGroup

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("协程 : %d\n", i)
			for j := 0; j < num2; j++ {
				atomic.AddInt64(&cnt, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("计算结果cnt : %d\n", cnt)
	fmt.Println("程序结束")
}
