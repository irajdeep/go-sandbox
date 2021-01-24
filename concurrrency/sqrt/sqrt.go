package main

import (
	"fmt"
	"sync"
)

func sqrtWorker(chIn chan int, chOut chan int) {
	fmt.Println("sqrtWorker started")
	for i := range chIn {
		sqrt := i * i
		chOut <- sqrt
	}
	fmt.Println("sqrtWorker finished")
}

func main() {
	var wg sync.WaitGroup

	chIn := make(chan int)
	chOut := make(chan int)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			sqrtWorker(chIn, chOut)
			wg.Done()
		}()
	}

	go func() {
		chIn <- 2
		chIn <- 4
		close(chIn)
	}()

	go func() {
		wg.Wait()
		close(chOut)
	}()

	for sqrt := range chOut {
		fmt.Printf("Got sqrt: %d\n", sqrt)
	}
}
