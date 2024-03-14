package email

import (
	"net/smtp"
	"os"
	"stori/email-service/user"
)

type EmailService struct {
	smtpHost    string
	smtpPort    string
	senderEmail string
	password    string
}

func NewEmailService() *EmailService {
	return &EmailService{
		smtpHost:    os.Getenv("SMTP_HOST"),
		smtpPort:    os.Getenv("SMTP_PORT"),
		senderEmail: os.Getenv("SENDER_EMAIL"),
		password:    os.Getenv("SENDER_PASSWORD"),
	}
}

func (e *EmailService) SendEmail(userData user.UserData, subject string, templateName string) error {

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte("From: Stori Challenge <" + e.senderEmail + ">/n" + "Subject:" + subject + "\n" + mime + templateName)

	auth := smtp.PlainAuth("", e.senderEmail, e.password, e.smtpHost)
	println("Auth: ", auth)
	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.senderEmail, []string{userData.Email}, message)
	println("Email send")
	if err != nil {
		println("Error to send email:", err.Error())
		return err
	}
	println("Email sended successfully!!")
	return nil

}
