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

type Authorizer interface {
	Authorize() string
}

type Capturer interface {
	Capture() string
}

type PaymentProvider interface {
	Authorizer
	Capturer
}

type StripeProvuder struct {
	Stripe string
}

func (s StripeProvuder) Authorize() string {
	return fmt.Sprintf("authorize stripe - %s", s.Stripe)
}
func (s StripeProvuder) Capture() string {
	return fmt.Sprintf("capture stripe - %s", s.Stripe)
}
func (s StripeProvuder) String() string {
	return fmt.Sprintf("печатаем stripe - %s", s.Stripe)
}

type SeriesResult struct {
	Title string
}

func InspectSearchResult(value any) {
	iss, ok := value.(string)
	fmt.Println(iss, ok)

	switch val := value.(type) {
	case string:
		fmt.Printf("String - %s\n", val)
	case int:
		fmt.Printf("Integer - %d\n", val)
	case bool:
		fmt.Printf("Bool - %v\n", val)
	case SeriesResult:
		fmt.Printf("SeriesResult - %s\n", val.Title)
	default:
		fmt.Printf("Other type - %v(%T)\n", val, val)
	}
}

type ModerationError struct {
	ContentID string
	Reason    string
}

func (m *ModerationError) Error() string {
	return fmt.Sprintf("Error, ContentID - %s, Reason - %s", m.ContentID, m.Reason)
}

func CheckPoster(status string) error {
	if status == "approved" {
		return nil
	} else {
		return &ModerationError{ContentID: "CONTENT #23sd", Reason: "Reason of error"}
	}

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

	var pprovider PaymentProvider

	pprovider = StripeProvuder{Stripe: "STRIPE"}
	fmt.Println(pprovider)

	fmt.Println(pprovider.Authorize(), pprovider.Capture())

	str := "String"
	integer := 123
	boolean := false
	sres := SeriesResult{Title: "SeriesResult"}
	stringslice := []string{"string", "string1", "strin5", "string6"}
	InspectSearchResult(str)
	InspectSearchResult(integer)
	InspectSearchResult(boolean)
	InspectSearchResult(sres)
	InspectSearchResult(stringslice)

	err := CheckPoster("approved")
	if err != nil {
		fmt.Println("Ошибка: ", err)
		if pErr, ok := err.(*ModerationError); ok {
			fmt.Printf("ContendID: %s, Reason: %s\n", pErr.ContentID, pErr.Reason)
		}
	}

	err = CheckPoster("rejected")
	if err != nil {
		fmt.Println("Ошибка: ", err)
		if pErr, ok := err.(*ModerationError); ok {
			fmt.Printf("ContendID: %s, Reason: %s\n", pErr.ContentID, pErr.Reason)
		}
	}

	var nilErr *ModerationError     // nil-pointer (empty)
	var errInterface error = nilErr // положили в интерфейс error

	fmt.Println("errInterface == nil?", errInterface == nil) // false
	fmt.Printf("Type: %T, Value: %v \n", errInterface, errInterface)
}
