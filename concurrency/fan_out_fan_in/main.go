package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(time.Millisecond * 500) // имитируем работу
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Fan-out: запускаем несколько горутин
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Отправляем задания
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Fan-in: собираем результаты
	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}
}
