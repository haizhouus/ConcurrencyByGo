package main

import (
	"fmt"
)


type Peterson struct {
	flag [2]bool
	victim int
}


func NewPeterson () *Peterson {
	return new(Peterson)
}


func (p *Peterson) lock(me int) {

	p.flag[me] = true
	p.victim = me
	for p.flag[1-me] && p.victim == me {}
}

func (p *Peterson) unlock(me int) {

	p.flag[me] = false
}

var x int = 0
var c [2]bool

func keepIncrement(p *Peterson, i int) {

	for j := 0; j < 100; j++ {
		p.lock(i)
		defer p.unlock(i)
		temp := x
		fmt.Println(x)
		x = temp + 1
	}
	c[i] = true
}


func main() {

	pet := NewPeterson()

	for i := 0; i < 2; i++ {
		go keepIncrement(pet, i)
	}
	
	for i := 0; i< 2; i++ {
		for !c[i] {}
	}
}
