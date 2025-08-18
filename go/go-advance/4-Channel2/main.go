package main

import (
	"fmt"
	"sync"
)

/**
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/

func main() {

	const bufferSize = 10
	const totalItems = 100
	dataChannel := make(chan int, bufferSize)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < totalItems; i++ {
			dataChannel <- i
			fmt.Printf("发送通道>>>>> : %v 通道大小: %d \n", i, len(dataChannel))
		}
	}()
	go func() {
		defer wg.Done()
		defer close(dataChannel)
		for data := range dataChannel {
			fmt.Printf("消费通道<<<<<<<<< : %v  通道大小 : %d  \n", data, len(dataChannel))
		}
	}()

	wg.Wait()
	fmt.Printf("程序结束")
}
