package main  
import (  
    "fmt"
    "sync"
    )

var x  = 0  
var c [1000]bool
var lock sync.Mutex

func increment(i int) {
    lock.Lock()
    x = x + 1
    lock.Unlock()
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
