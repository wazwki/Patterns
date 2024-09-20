package main

import "fmt"

// Mediator - интерфейс посредника
type Mediator interface {
	SendMessage(sender string, message string)
}

// User - интерфейс пользователя
type User interface {
	Send(message string)
	Receive(sender string, message string)
}

// ChatRoom - конкретная реализация посредника (чат-комната)
type ChatRoom struct {
	users []User
}

func (c *ChatRoom) AddUser(user User) {
	c.users = append(c.users, user)
}

func (c *ChatRoom) SendMessage(sender string, message string) {
	for _, user := range c.users {
		user.Receive(sender, message)
	}
}

// ConcreteUser - конкретный пользователь чата
type ConcreteUser struct {
	name     string
	mediator Mediator
}

func (u *ConcreteUser) Send(message string) {
	fmt.Printf("%s отправляет сообщение: %s\n", u.name, message)
	u.mediator.SendMessage(u.name, message)
}

func (u *ConcreteUser) Receive(sender string, message string) {
	if sender != u.name {
		fmt.Printf("%s получает сообщение от %s: %s\n", u.name, sender, message)
	}
}

func main() {
	// Создаем чат-комнату (посредник)
	chatRoom := &ChatRoom{}

	// Создаем пользователей
	user1 := &ConcreteUser{name: "Alice", mediator: chatRoom}
	user2 := &ConcreteUser{name: "Bob", mediator: chatRoom}
	user3 := &ConcreteUser{name: "Charlie", mediator: chatRoom}

	// Добавляем пользователей в чат-комнату
	chatRoom.AddUser(user1)
	chatRoom.AddUser(user2)
	chatRoom.AddUser(user3)

	// Пользователи отправляют сообщения через посредника (чат-комнату)
	user1.Send("Привет всем!")
	user2.Send("Привет, Alice!")
	user3.Send("Привет, Bob!")
}
