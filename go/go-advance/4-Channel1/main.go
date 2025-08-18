package main

import (
	"fmt"
	"time"
)

/*
*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func main() {

	dataChannel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			dataChannel <- i
			fmt.Println("发送通道:", i)
			time.Sleep(100 * time.Millisecond) // 模拟处理时间
		}
		close(dataChannel) // 发送完成后关闭通道
		fmt.Println("生产者完成")
	}()
	go func() {
		for {
			num, ok := <-dataChannel
			if !ok {
				println("消费者结束")
				return
			}
			println("消费通道:", num)
		}
		time.Sleep(150 * time.Millisecond)
	}()
	time.Sleep(2 * time.Second) // 主协程等待足够时间
	fmt.Println("程序结束")
}
