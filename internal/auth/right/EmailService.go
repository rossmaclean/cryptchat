package authright

import (
	"cryptchat/internal"
	"log"
	"net/smtp"
	"os"
)

var smtpServer string
var smtpPort string
var smtpsPort string
var smtpUsername string
var smtpPassword string
var fromAddr string
var domain string

func init() {
	if !internal.IsLocal() {
		smtpServer = os.Getenv("CLOUDRON_MAIL_SMTP_SERVER")
		smtpPort = os.Getenv("CLOUDRON_MAIL_SMTP_PORT")
		smtpsPort = os.Getenv("CLOUDRON_MAIL_SMTPS_PORT")
		smtpUsername = os.Getenv("CLOUDRON_MAIL_SMTP_USERNAME")
		smtpPassword = os.Getenv("CLOUDRON_MAIL_SMTP_PASSWORD")
		fromAddr = os.Getenv("CLOUDRON_MAIL_FROM")
		domain = os.Getenv("CLOUDRON_MAIL_DOMAIN")

		log.Println("Sending mail")
		SendMail()
	}
}

func SendMail() error {
	to := []string{"ross@rossmac.co.uk"}
	message := []byte("My super secret message.")

	auth := smtp.PlainAuth("", fromAddr, smtpPassword, smtpServer)

	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, fromAddr, to, message)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
