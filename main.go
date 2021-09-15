package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"questionsandanswers/repository"
	"questionsandanswers/services"

	"questionsandanswers/services/transport"
	httpTransport "questionsandanswers/services/transport/http"
	"syscall"
)

func main() {

	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)

	flag.Parse()

	log.Println("Service started")
	defer log.Println("Service ended")

	dbclient, err := repository.GetMontoDbClient()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(dbclient.Connect(context.Background()))

	rep := repository.MongoDbRepository{}
	s := services.NewService(rep)
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := transport.MakeEndpoints(s)

	go func() {
		log.Println("Listening on port: ", *httpAddr)
		handler := httpTransport.NewService(endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Default().Fatalln(<-errChan)

}
