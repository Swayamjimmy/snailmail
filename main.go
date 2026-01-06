package main

import "time"

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	go loadRecipient("emails.csv", recipientChannel)

	for workers := 0; workers <= 4; workers++ {
		go emailWorker(workers, recipientChannel)
	}

	time.Sleep(3 * time.Second)
}
