package mailer

import (
	"fmt"
	"net/smtp"
)

// Config ...
type Config struct {
	Mail    *string
	Email   string
	Pending int
	Opt     *SMTPOption
}

// SMTPOption ...
type SMTPOption struct {
	From        string
	Pass        string
	SMTPAddress string
	SMTPPort    string
}

// SendMail ...
func (mail Config) SendMail() bool {
	from := mail.Opt.From
	pass := mail.Opt.Pass
	port := mail.Opt.SMTPPort
	address := mail.Opt.SMTPAddress

	smtpServer := fmt.Sprintf("%s:%s", address, port)
	err := smtp.SendMail(smtpServer, smtp.PlainAuth("", from, pass, address), from, []string{mail.Email}, []byte(*mail.Mail))

	fmt.Println("Sent Mail " + mail.Email)

	return err == nil
}

// FilterMails ...
func FilterMails(mails *[]Config) []string {
	var result []string

	for i := range *mails {
		result = append(result, (*mails)[i].Email)
	}

	return result
}
