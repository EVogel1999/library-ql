package main

import (
	"github.com/graphql-go/graphql"
	"library-ql/controllers"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"author": controllers.FindAuthorById,
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})
