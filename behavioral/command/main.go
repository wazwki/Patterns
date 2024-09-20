package main

import "fmt"

// Command - интерфейс для команд
type Command interface {
	Execute()
	Undo()
}

// Light - получатель команды (лампа)
type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("The light is on")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("The light is off")
}

// LightOnCommand - команда для включения света
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

func (c *LightOnCommand) Undo() {
	c.light.TurnOff()
}

// LightOffCommand - команда для выключения света
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

func (c *LightOffCommand) Undo() {
	c.light.TurnOn()
}

// RemoteControl - отправитель команд
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func (r *RemoteControl) PressUndo() {
	r.command.Undo()
}

func main() {
	// Создаем объект света (получатель)
	light := &Light{}

	// Создаем команды для включения и выключения света
	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	// Создаем пульт дистанционного управления
	remote := &RemoteControl{}

	// Включаем свет
	remote.SetCommand(lightOn)
	remote.PressButton() // Output: The light is on

	// Выключаем свет
	remote.SetCommand(lightOff)
	remote.PressButton() // Output: The light is off

	// Отменяем выключение света
	remote.PressUndo() // Output: The light is on
}
