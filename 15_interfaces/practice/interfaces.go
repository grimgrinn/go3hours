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

type Exporter interface {
	Export() string
}

type CSVExpot struct {
	Field string
}

func (c CSVExpot) Export() string {
	return fmt.Sprintf("CSVExport %s", c.Field)
}

type JSONExport struct {
	Field string
}

func (j JSONExport) Export() string {
	return fmt.Sprintf("JSONExxport %s", j.Field)
}

func RunExport(e Exporter) {
	fmt.Println(e.Export())
}

func main() {
	emailn := EmailNotification{Message: "email"}
	pushn := PushNotification{Message: "push"}

	notifiers := []Notifier{emailn, pushn}

	for _, n := range notifiers {
		fmt.Println(n.ChannelName(), n.Send())
	}

	cex := CSVExpot{Field: "CSV"}
	jex := JSONExport{Field: "JSON"}

	RunExport(cex)
	RunExport(jex)
}
