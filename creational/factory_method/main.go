package main

import "fmt"

// NotificationFactory
type NotificationFactory interface {
	setText(string)
	getText() string
}

// Notification
type Notification struct {
	Text string
}

func (n *Notification) setText(text string) {
	n.Text = text
}

func (n *Notification) getText() string {
	return n.Text
}

// email
type Email struct {
	Notification
}

func getEmail() NotificationFactory {
	return &Email{
		Notification{
			Text: "Email",
		},
	}
}

// SMS
type SMS struct {
	Notification
}

func getSMS() NotificationFactory {
	return &SMS{
		Notification{
			Text: "SMS",
		},
	}
}

// push
type Push struct {
	Notification
}

func getPush() NotificationFactory {
	return &Push{
		Notification{
			Text: "Push",
		},
	}
}

// messager
type Messager struct {
	Notification
}

func getMessager() NotificationFactory {
	return &Messager{
		Notification{
			Text: "Messager",
		},
	}
}

func getNotification(nType string) NotificationFactory {
	switch nType {
	case "email":
		return getEmail()
	case "sms":
		return getSMS()
	case "push":
		return getPush()
	case "messager":
		return getMessager()
	}
	return nil
}

func main() {
	email := getNotification("email")
	fmt.Println(email.getText())

	sms := getNotification("sms")
	fmt.Println(sms.getText())

	push := getNotification("push")
	fmt.Println(push.getText())

	messager := getNotification("messager")
	fmt.Println(messager.getText())
}
