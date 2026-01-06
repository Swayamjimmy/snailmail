package main

import "fmt"

func emailWorker(id int, ch chan Recipient) {
	for recipient := range ch {
		fmt.Printf("Worker %d sending email to %s\n", id, recipient.Email)
	}
}
