package main

import (
	"fmt"
	"time"
)

func stage1(nums []int, out chan<- int) {
	for _, n := range nums {
		out <- n
	}
	close(out)
}

func stage2(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * 2
	}
	close(out)
}

func stage3(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Создаем каналы для передачи данных между стадиями
	stage1Chan := make(chan int)
	stage2Chan := make(chan int)

	// Запускаем стадии
	go stage1(numbers, stage1Chan)
	go stage2(stage1Chan, stage2Chan)
	go stage3(stage2Chan)

	time.Sleep(time.Second) // даем время на завершение всех стадий
}
