package authright

import (
	"cryptchat/internal"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"os"
	"strconv"
)

var smtpServer string
var smtpPort int
var smtpsPort string
var smtpUsername string
var smtpPassword string
var fromAddr string
var domain string

func init() {
	if !internal.IsLocal() {
		smtpServer = os.Getenv("CLOUDRON_MAIL_SMTP_SERVER")
		var err error
		smtpPort, err = strconv.Atoi(os.Getenv("CLOUDRON_MAIL_SMTP_PORT"))
		if err != nil {
			log.Fatal(err)
		}
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
	server.Host = smtpServer
	server.Port = smtpPort
	server.Username = smtpUsername
	server.Password = smtpPassword
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Printf("Mail error: %s", err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("ross@rossmac.co.uk")
	email.AddTo("ross@rossmac.co.uk")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
