package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"
	"library-ql/controllers"
	"library-ql/database"
	"log"
	"net/http"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"author": controllers.FindAuthorById,
		"book":   controllers.FindBookById,
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error configuring .env file")
	}

	database.Connect()
	initGraphQL()
}

func initGraphQL() {
	handler := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil)
}
