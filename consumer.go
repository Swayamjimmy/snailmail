package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipient := range ch {

		smtpHost := "localhost"
		smtpPort := "1025"

		formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s", recipient.Email, "Hello, this is a test email from Snailmail")

		msg := []byte(formattedMsg)

		err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "snailmail@example.com", []string{recipient.Email}, msg)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d sent email to %s\n", id, recipient.Email)

	}
}
