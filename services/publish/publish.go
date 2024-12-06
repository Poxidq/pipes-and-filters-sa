package publish

import (
	"log"
	"paf/config"
	"paf/shared"

	gomail "gopkg.in/mail.v2"
)

type Publish struct {
	Next shared.Service
}

func (c *Publish) Execute(p *shared.Process) {
	sendEmail("Test subject", p.Message.Message)
	log.Printf("Sended message!")
	if c.Next != nil {
		c.Next.Execute(p)
	}
}

func (c *Publish) SetNext(next shared.Service) {
	c.Next = next
}

func sendEmail(subject, body string) error {
	emailRecipients := config.CFG.EmailRecipients
	fromEmail := config.CFG.EmailAddress
	emailPassword := config.CFG.EmailPassword

	m := gomail.NewMessage()

	m.SetHeader("From", fromEmail)
	m.SetHeader("To", emailRecipients)
	m.SetHeader("Subject", subject)

	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, fromEmail, emailPassword)

	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		panic(err)
	}
	return nil
}
