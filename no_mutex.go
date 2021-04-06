package main  
import (  
    "fmt"
    )

var x  = 0  
var c [1000]bool

func increment(i int) {  
    x = x + 1
    c[i] = true
}

func main() {  

    for i := 0; i < 1000; i++ {
         go increment(i)
    }
    for i := 0; i < 1000; i++ {
	for !c[i] {}
    }
    fmt.Println("final value of x =", x)
}
