package main

import (
	"log"

	"github.com/AlexDespod/testtask/pkg"
)

func main() {
	msgs, conn, ch := pkg.GetConsumer("hello")
	defer conn.Close()
	defer ch.Close()

	forever := make(chan struct{})

	var pathToFile string

	go func() {

		for d := range msgs {

			pathToFile = string(d.Body)

			go pkg.Resize(pathToFile)

			log.Printf("Received a message: %s", d.Body)
		}

	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
