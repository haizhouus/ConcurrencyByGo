package main

import (
	"fmt"
)

var x, y [100]int
var c, d int

func operation(i int) {
	for d == 0 {}
	x[i] = 1
	if i > 0 {
		y[i] = x[i-1]
	} else {
		y[i] = x[99]
	}
	temp := c
	temp++
	temp--
	temp++
	c = temp
}

func main() {
	fmt.Println("Hello, Concurrency!")
	for i := 0; i < 100; i++ {
		go operation(i)
	}
	d = 1
	for c < 100 {}
	for i := 0; i < 100; i++ {
		fmt.Println(x[i], y[i])
	}
}
