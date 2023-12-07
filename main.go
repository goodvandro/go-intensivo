package main

import (
	"fmt"
	"time"
)

// T1
func main() {
	canal := make(chan int) // make a chanel

	go func() {
		for i := 0; i < 10; i++ {
			canal <- i // set data into the chanel
			fmt.Println("Jogou no canal", i)
		}
	}()

	go worker(canal, 1)
	worker(canal, 2)

}

func worker(canal chan int, workerID int) {
	for {
		fmt.Println("Recebeu do canal", <-canal, "no worker", workerID)
		time.Sleep(time.Second)
	}
}
