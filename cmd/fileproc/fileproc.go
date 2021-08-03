//fileproc is program that launch consumer , which does resizing of photos
package main

import (
	"log"
	"runtime"

	"github.com/AlexDespod/testtask/pkg"
	"github.com/AlexDespod/testtask/shared"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	consumer, err := pkg.GetConsumer(shared.QueueName)

	if err != nil {
		panic(err)
	}

	defer consumer.CloseAll()

	forever := make(chan struct{})

	var pathToFile string

	go func() {

		for d := range consumer.Delivery {

			pathToFile = string(d.Body)

			go pkg.Resize(pathToFile, 200, 0)

			log.Printf("Received a message: %s", pathToFile)
		}

	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
