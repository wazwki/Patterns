package main

import (
	"fmt"
	"time"
)

func asyncTask() chan int {
	result := make(chan int)
	go func() {
		time.Sleep(2 * time.Second) // Имитируем долгую работу
		result <- 42
	}()
	return result
}

func main() {
	resultChan := asyncTask()
	fmt.Println("Doing other work...")
	result := <-resultChan
	fmt.Println("Result:", result)
}
