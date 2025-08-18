package main

import "fmt"

func main() {
	/**
	指针 1
	*/
	a := 1
	cnt(&a)
	fmt.Println(a)
}

/*
*
指针 1
*/
func cnt(a *int) {
	*a++
}
