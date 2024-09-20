package main

import (
	"fmt"
	"time"
)

func main() {
	rate := time.Tick(500 * time.Millisecond) // Лимит 1 операция каждые 500 мс
	for i := 1; i <= 5; i++ {
		<-rate // Ждем сигнала от тикера
		fmt.Println("Operation", i, "executed at", time.Now())
	}
}
