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

	conn, publishChan, queue := pkg.GetPublisher()
	defer conn.Close()
	defer publishChan.Close()

	publisher := shared.Publisher{Chan: publishChan, Queue: queue}

	http.HandleFunc("/", injectConn(publisher, httphandl))

	log.Fatal(http.ListenAndServe(":9090", nil))
}

func injectConn(publisher shared.Publisher, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "publisher", publisher)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func httphandl(res http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	uploadedFile, header, err := req.FormFile("file")

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(500)
		return
	}

	path := pkg.GetName(header.Filename)

	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(500)
		return
	}

	defer file.Close()

	io.Copy(file, uploadedFile)

	pkg.SendMessToMQ(req.Context(), path)

	res.WriteHeader(200)
}
