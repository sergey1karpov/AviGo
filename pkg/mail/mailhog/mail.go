package mailhog

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(to string) {
	smtpHost := "localhost"
	smtpPort := 1025
	smtpAddr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	from := "avigo@support.com"

	subject := "Привет, мир!"
	body := "Это тело письма."

	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", "", "", smtpHost)

	err := smtp.SendMail(smtpAddr, auth, from, []string{to}, message)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Письмо успешно отправлено!")
	}
}
