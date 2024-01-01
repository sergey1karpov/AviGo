package mailhog

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"time"
)

func SendMail(to, template string) {
	smtpHost := "localhost"
	smtpPort := 1025
	smtpAddr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	from := "avigo@support.com"

	subject := "Привет, мир!"

	htmlBody := strings.ReplaceAll(strings.ReplaceAll(template, "{LINK}", "<a href=\"/login\">Login Page</a>"), "{YEAR}", fmt.Sprintf("%d", time.Now().Year()))

	message := []byte("To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		htmlBody + "\r\n")

	auth := smtp.PlainAuth("", "", "", smtpHost)

	err := smtp.SendMail(smtpAddr, auth, from, []string{to}, message)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Письмо успешно отправлено!")
	}
}
