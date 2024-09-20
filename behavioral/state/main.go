package main

import "fmt"

// State - интерфейс состояния
type State interface {
	PressButton(fan *Fan)
}

// OffState - состояние "выключено"
type OffState struct{}

func (s *OffState) PressButton(fan *Fan) {
	fmt.Println("Fan is turning on at low speed.")
	fan.SetState(&LowSpeedState{})
}

// LowSpeedState - состояние "низкая скорость"
type LowSpeedState struct{}

func (s *LowSpeedState) PressButton(fan *Fan) {
	fmt.Println("Fan is turning to high speed.")
	fan.SetState(&HighSpeedState{})
}

// HighSpeedState - состояние "высокая скорость"
type HighSpeedState struct{}

func (s *HighSpeedState) PressButton(fan *Fan) {
	fmt.Println("Fan is turning off.")
	fan.SetState(&OffState{})
}

// Fan - контекст, который изменяет поведение в зависимости от состояния
type Fan struct {
	currentState State
}

func (f *Fan) SetState(state State) {
	f.currentState = state
}

func (f *Fan) PressButton() {
	f.currentState.PressButton(f)
}

func main() {
	// Создаем вентилятор и устанавливаем начальное состояние (выключен)
	fan := &Fan{currentState: &OffState{}}

	// Нажимаем на кнопку несколько раз, чтобы переключать состояния
	fan.PressButton() // Включение на низкой скорости
	fan.PressButton() // Включение на высокой скорости
	fan.PressButton() // Выключение
	fan.PressButton() // Включение на низкой скорости
}
