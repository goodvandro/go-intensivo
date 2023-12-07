package main

import (
	"fmt"
	"time"
)

func processando() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// T1
func main() {
	canal := make(chan int) // make a chanel

	// set data into the chanel
	go func() {
		canal <- 1 // T2
	}()

	// remove data from the chanel
	fmt.Println(<-canal)

	time.Sleep(time.Second * 2)
}
