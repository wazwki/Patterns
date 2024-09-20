package main

import "fmt"

// Beverage - шаблонный метод для приготовления напитков
type Beverage interface {
	BoilWater()
	Brew() // Шаг, который будет отличаться для разных напитков
	PourInCup()
	AddCondiments() // Шаг, который будет отличаться для разных напитков
	Make()
}

// BaseBeverage - базовая структура с общими шагами для приготовления напитков
type BaseBeverage struct{}

func (b *BaseBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

func (b *BaseBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}

// Coffee - конкретная реализация для кофе
type Coffee struct {
	BaseBeverage
}

func (c *Coffee) Brew() {
	fmt.Println("Brewing coffee")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk to coffee")
}

func (c *Coffee) Make() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()
}

// Tea - конкретная реализация для чая
type Tea struct {
	BaseBeverage
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon to tea")
}

func (t *Tea) Make() {
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	t.AddCondiments()
}

func main() {
	// Приготовление кофе
	coffee := &Coffee{}
	fmt.Println("Making coffee:")
	coffee.Make()

	// Приготовление чая
	tea := &Tea{}
	fmt.Println("\nMaking tea:")
	tea.Make()
}
