package main

import (
	"github.com/graphql-go/graphql"
	"library-ql/author"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"author": author.FindAuthorById,
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})
