//fileproc is program that launch simple http server , which receive files and comunicate with RabbitMQ
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/AlexDespod/testtask/pkg"
	"github.com/AlexDespod/testtask/shared"
)

func main() {

	publisher, err := pkg.GetPublisher(shared.QueueName)

	if err != nil {
		panic(err)
	}

	defer publisher.CloseAll()

	http.HandleFunc("/", injectConn(publisher, httphandl))

	log.Fatal(http.ListenAndServe(":9090", nil))
}

//injectConn pull rabbitMQ connection to all requests
func injectConn(publisher *pkg.Publisher, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "publisher", publisher)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

//httphandl is function wich handle a requests
func httphandl(res http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	uploadedFile, header, err := req.FormFile("file")

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	defer uploadedFile.Close()

	path := pkg.GetName(header.Filename)

	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(500)
		return
	}

	defer file.Close()

	io.Copy(file, uploadedFile)

	err = pkg.SendMessToMQ(req.Context(), path)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(500)
		return
	}

	res.WriteHeader(200)
}
