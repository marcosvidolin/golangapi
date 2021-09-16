package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"questionsandanswers/repository"
	"questionsandanswers/router"
)

func main() {

	dbclient, err := repository.GetMontoDbClient()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dbclient.Connect(context.Background()))

	router := router.HttpHandler()
	log.Fatal(http.ListenAndServe(":8080", router))
}
