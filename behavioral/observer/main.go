package main

import "fmt"

// Observer - интерфейс наблюдателя
type Observer interface {
	Update(stockPrice float64)
}

// StockSubject - интерфейс субъекта (объект, за которым наблюдают)
type StockSubject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify()
}

// Stock - конкретная реализация субъекта (акция)
type Stock struct {
	observers  []Observer
	stockPrice float64
}

func (s *Stock) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Stock) Unregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Stock) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.stockPrice)
	}
}

func (s *Stock) SetPrice(price float64) {
	fmt.Printf("Setting stock price to %.2f\n", price)
	s.stockPrice = price
	s.Notify()
}

// Investor - конкретная реализация наблюдателя (инвестор)
type Investor struct {
	name string
}

func (i *Investor) Update(stockPrice float64) {
	fmt.Printf("Investor %s has been notified of new stock price: %.2f\n", i.name, stockPrice)
}

func main() {
	// Создаем объект акции
	stock := &Stock{}

	// Создаем инвесторов
	investor1 := &Investor{name: "Alice"}
	investor2 := &Investor{name: "Bob"}

	// Регистрируем инвесторов
	stock.Register(investor1)
	stock.Register(investor2)

	// Устанавливаем цену акции и уведомляем всех наблюдателей
	stock.SetPrice(100.50)
	stock.SetPrice(105.75)

	// Убираем одного из наблюдателей
	stock.Unregister(investor1)

	// Изменяем цену акции и уведомляем оставшихся наблюдателей
	stock.SetPrice(110.00)
}
