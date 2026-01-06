package main

import (
	"fmt"
	"sync"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipient := range ch {
		fmt.Printf("Worker %d sending email to %s\n", id, recipient.Email)
	}
}
