package main

import "fmt"

type Notifier interface {
	Send() string
	ChannelName() string
}

type EmailNotification struct {
	Message string
}

func (e EmailNotification) Send() string {
	return fmt.Sprintf("email was sent, message - %s\n", e.Message)
}
func (e EmailNotification) ChannelName() string {
	return "email\n"
}

type PushNotification struct {
	Message string
}

func (p PushNotification) Send() string {
	return fmt.Sprintf("push was sent, message - %s\n", p.Message)
}
func (p PushNotification) ChannelName() string {
	return "push\n"
}

func main() {
	emailn := EmailNotification{Message: "email"}
	pushn := PushNotification{Message: "push"}

	notifiers := []Notifier{emailn, pushn}

	for _, n := range notifiers {
		fmt.Println(n.ChannelName(), n.Send())
	}
}
