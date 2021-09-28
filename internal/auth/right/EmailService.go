package authright

import (
	"cryptchat/internal"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
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

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, World</title>
</head>
<body>
   <p>This is an email using Go</p>
</body>
`

func SendMail() error {
	log.Printf("Server: %s", smtpServer)
	log.Printf("Domain: %s", domain)
	server := mail.NewSMTPClient()
	server.Host = "smtp.host.com"
	server.Port = 587
	server.Username = "username@host.com"
	server.Password = "supersecretpassword"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("From Me <me@host.com>")
	email.AddTo("you@example.com")
	email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)
	email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
}
