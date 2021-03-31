package main

import (
	"fmt"
)

var x, y [100]int

func operation(i int) {
	x[i] = 1
	if i > 0 {
		y[i] = x[i-1]
	} else {
		y[i] = x[99]
	}
}
		
func main() {
	fmt.Println("Hello, Concurrency!")
	for i := 0; i < 100; i++ {
		go operation(i)
	}
	for i :=0; i < 100; i++ {
		fmt.Println(x[i], y[i])
	}	
}
