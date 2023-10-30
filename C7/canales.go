package main

import (
	"fmt"
)

func main() {
	n := 3

	ch := make(chan int)

	go multipPorDos(n, ch)

	fmt.Println(<-ch)
}

func multipPorDos(num int, ch chan int) {
	res := num * 2
	ch <- res
}