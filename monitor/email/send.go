package email

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Send(subject string, content string) {
	from := mail.NewEmail(os.Getenv("SENDGRID_FROM_NAME"), os.Getenv("SENDGRID_FROM"))
	to := mail.NewEmail(os.Getenv("SENDGRID_TO_NAME"), os.Getenv("SENDGRID_TO"))
	message := mail.NewSingleEmail(from, subject, to, content, "")
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_KEY"))
	_, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
