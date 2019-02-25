package mailer

import (
	"fmt"
	"net/smtp"
)

type MailerConfig struct {
	Mail  *string
	Email string
	Fails int
	Id    int
}

func (mail MailerConfig) SendMail() bool {
	/*
		// Concurrent testing purposes only
		rand.Seed(time.Now().UTC().UnixNano())
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	*/

	from := "...@gmail.com"
	pass := "..."

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{mail.Email}, []byte(*mail.Mail))

	fmt.Println("SentMail " + mail.Email)

	return err == nil
}

type MailerApi struct {
	Get string
	Put string
}

func (api MailerApi) putUrl(mail MailerConfig) string {
	return fmt.Sprintf(api.Put, mail.Id)
}

func (api MailerApi) Update(mail MailerConfig) {
	fmt.Println("Updated " + mail.Email)
	// Submit to the API the mail with number of failures plus one
}
