package main

import (
	"fmt"
	"sync"
)

func main() {

	/**
	Goroutine 1
	*/
	cnt := 1
	n := 10
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			if cnt > n {
				mu.Unlock()
				return
			}
			if cnt%2 != 0 {
				fmt.Println("奇数:", cnt)
				cnt++
			}
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			if cnt > n {
				mu.Unlock()
				return
			}
			if cnt%2 == 0 {
				fmt.Println("偶数:", cnt)
				cnt++
			}
			mu.Unlock()
		}
	}()
	wg.Wait()
}
