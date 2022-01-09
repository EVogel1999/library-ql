package main

import (
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"library-ql/author"
	"library-ql/books"
	"library-ql/database"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error configuring .env file")
	}

	database.Connect()
	initControllers()
	initGraphQL()
}

func initControllers() {
	author.InitController()
	books.InitController()
}

func initGraphQL() {
	handler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil)
}
