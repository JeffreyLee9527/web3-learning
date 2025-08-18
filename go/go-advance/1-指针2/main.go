package main

import "fmt"

func main() {
	/**
	指针 2
	*/
	slicePtr := []int{1, 2, 3}
	fmt.Println(slicePtr)
	doubleSlice(&slicePtr)
	fmt.Println(slicePtr)
}

/*
*
指针 2
*/
func doubleSlice(slicePtr *[]int) {
	if slicePtr == nil {
		return
	}
	slice := *slicePtr //还原切片
	for i := range slice {
		slice[i] *= 2
	}
}
