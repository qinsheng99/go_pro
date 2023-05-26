package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	var (
		a = make(chan struct{}, 1)
		b = make(chan struct{})
		c = make(chan struct{})
	)
	a <- struct{}{}
	go A(a, b)
	go B(b, c)
	go C(c)

	wg.Wait()
}

func A(a chan struct{}, b chan struct{}) {
	defer wg.Done()
	<-a
	fmt.Println("A")
	b <- struct{}{}

}

func B(b chan struct{}, c chan struct{}) {
	defer wg.Done()
	<-b
	fmt.Println("B")
	c <- struct{}{}
}

func C(c chan struct{}) {
	defer wg.Done()
	<-c
	fmt.Println("C")
}
