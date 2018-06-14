package main

import (
	"fmt"
	"runtime"
)

func main1() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}

// Go nm comment
func Go(c chan bool, index int) {
	a := 1000000000
	fmt.Println(a, index)
	c <- true
}
