package main

import (
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	go loadRecipient("emails.csv", recipientChannel)

	var wg sync.WaitGroup

	for workers := 0; workers <= 4; workers++ {
		wg.Add(1)
		go emailWorker(workers, recipientChannel, &wg)
	}

	wg.Wait()

}
