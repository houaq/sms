package lib

import (
	"log"
	"time"

	"github.com/houaq/sms/modem"
)

// InitWorker starts the worker
func InitWorker() {
	messages := make(chan SMS)
	go producer(messages)
	go consumer(messages)
}

func consumer(messages chan SMS) {
	for {
		message := <-messages
		log.Println("consumer: processing", message.UUID)
		err := modem.SendMessage(message.Mobile, message.Body)
		if err != nil {
			message.Status = "error"
			log.Println("consumer: failed to process", message.UUID, err)
		} else {
			message.Status = "sent"
		}
		message.Retries++
		// TODO: make this update a goroutine?
		UpdateMessageStatus(message)
	}
}

func producer(messages chan SMS) {
	for {
		pendingMsgs, err := GetPendingMessages()
		if err != nil {
			log.Printf("producer: failed to get messages. %s", err.Error())
		}
		// log.Printf("producer: %d pending messages found", len(pendingMsgs))
		for _, msg := range pendingMsgs {
			log.Printf("producer: Processing %#v", msg)
			messages <- msg
		}
		time.Sleep(time.Second)
	}
}
