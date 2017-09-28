package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func NewIPhone() *IPhone {
	return &IPhone{}
}
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func x(args Phone) {
	args.call()
}
func main() {
	x(NewIPhone())
	x(new(NokiaPhone))
}
