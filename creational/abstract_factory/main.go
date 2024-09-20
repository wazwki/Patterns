package main

import "fmt"

// Абстрактная фабрика
type OsFactory interface {
	createButton() IButton
	createTextField() ITextField
	createWindow() IWindow
}

// Интерфейсы для компонентов
type IButton interface {
	setSize(int)
	getSize() int
}

type ITextField interface {
	setSize(int)
	getSize() int
}

type IWindow interface {
	setSize(int)
	getSize() int
}

// Реализация компонентов
type Button struct {
	size int
}

type TextField struct {
	size int
}

type Window struct {
	size int
}

func (b *Button) setSize(size int) {
	b.size = size
}

func (b *Button) getSize() int {
	return b.size
}

func (t *TextField) setSize(size int) {
	t.size = size
}

func (t *TextField) getSize() int {
	return t.size
}

func (w *Window) setSize(size int) {
	w.size = size
}

func (w *Window) getSize() int {
	return w.size
}

// Реализация для Windows
type Windows struct{}

func (w *Windows) createButton() IButton {
	return &Button{}
}

func (w *Windows) createTextField() ITextField {
	return &TextField{}
}

func (w *Windows) createWindow() IWindow {
	return &Window{}
}

// Реализация для MacOS
type MacOS struct{}

func (m *MacOS) createButton() IButton {
	return &Button{}
}

func (m *MacOS) createTextField() ITextField {
	return &TextField{}
}

func (m *MacOS) createWindow() IWindow {
	return &Window{}
}

// Реализация для Linux
type Linux struct{}

func (l *Linux) createButton() IButton {
	return &Button{}
}

func (l *Linux) createTextField() ITextField {
	return &TextField{}
}

func (l *Linux) createWindow() IWindow {
	return &Window{}
}

// Использование фабрики
func main() {
	// Windows
	windowsFactory := &Windows{}
	winButton := windowsFactory.createButton()
	winTextField := windowsFactory.createTextField()
	winWindow := windowsFactory.createWindow()

	winButton.setSize(10)
	winTextField.setSize(10)
	winWindow.setSize(10)

	fmt.Println("Windows Button size:", winButton.getSize())
	fmt.Println("Windows TextField size:", winTextField.getSize())
	fmt.Println("Windows Window size:", winWindow.getSize())

	// MacOS
	macFactory := &MacOS{}
	macButton := macFactory.createButton()
	macTextField := macFactory.createTextField()
	macWindow := macFactory.createWindow()

	macButton.setSize(12)
	macTextField.setSize(12)
	macWindow.setSize(12)

	fmt.Println("MacOS Button size:", macButton.getSize())
	fmt.Println("MacOS TextField size:", macTextField.getSize())
	fmt.Println("MacOS Window size:", macWindow.getSize())

	// Linux
	linuxFactory := &Linux{}
	linButton := linuxFactory.createButton()
	linTextField := linuxFactory.createTextField()
	linWindow := linuxFactory.createWindow()

	linButton.setSize(14)
	linTextField.setSize(14)
	linWindow.setSize(14)

	fmt.Println("Linux Button size:", linButton.getSize())
	fmt.Println("Linux TextField size:", linTextField.getSize())
	fmt.Println("Linux Window size:", linWindow.getSize())
}
