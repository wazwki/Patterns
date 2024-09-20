package main

import "fmt"

// Shape - интерфейс для фигур
type Shape interface {
	Draw()
}

// Circle - конкретная реализация фигуры "Круг"
type Circle struct {
	color  string
	x, y   int
	radius int
}

func (c *Circle) SetX(x int) {
	c.x = x
}

func (c *Circle) SetY(y int) {
	c.y = y
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing Circle: Color = %s, X = %d, Y = %d, Radius = %d\n", c.color, c.x, c.y, c.radius)
}

// CircleFactory - фабрика, использующая паттерн Flyweight для повторного использования кругов одного цвета
type CircleFactory struct {
	circleMap map[string]*Circle
}

func NewCircleFactory() *CircleFactory {
	return &CircleFactory{
		circleMap: make(map[string]*Circle),
	}
}

func (cf *CircleFactory) GetCircle(color string) *Circle {
	if cf.circleMap[color] == nil {
		cf.circleMap[color] = &Circle{color: color}
		fmt.Printf("Creating new circle with color: %s\n", color)
	}
	return cf.circleMap[color]
}

func main() {
	circleFactory := NewCircleFactory()

	// Создаем круги с использованием фабрики
	redCircle := circleFactory.GetCircle("Red")
	redCircle.SetX(10)
	redCircle.SetY(20)
	redCircle.SetRadius(15)
	redCircle.Draw()

	blueCircle := circleFactory.GetCircle("Blue")
	blueCircle.SetX(30)
	blueCircle.SetY(40)
	blueCircle.SetRadius(20)
	blueCircle.Draw()

	// Используем существующий "Red" круг повторно
	anotherRedCircle := circleFactory.GetCircle("Red")
	anotherRedCircle.SetX(50)
	anotherRedCircle.SetY(60)
	anotherRedCircle.SetRadius(25)
	anotherRedCircle.Draw()

	// Заметьте, что для второго красного круга фабрика не создает новый объект
}
