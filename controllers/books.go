package controllers

import (
	"errors"
	"github.com/graphql-go/graphql"
	"library-ql/database"
)

// This is slightly different from the one in author.go because the other one would cause a loop
var bookAuthorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BookAuthor",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"pseudonym": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"isbn": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"download": &graphql.Field{
			Type: graphql.String,
		},
		"authors": &graphql.Field{
			Type: graphql.NewList(bookAuthorType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorIDs := p.Source.(database.Book).Authors
				if len(authorIDs) > 0 {
					if authors, err := database.GetAuthorsByIDs(authorIDs); err != nil {
						return nil, err
					} else {
						return authors, nil
					}
				}
				return nil, errors.New("could not parse author ids for book")
			},
		},
	},
})

var FindBookById = &graphql.Field{
	Type:        bookType,
	Description: "Gets a single book",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(string)
		if ok {
			if book, err := database.GetBookByID(id); err != nil {
				return nil, err
			} else {
				return book, nil
			}
		}
		return nil, errors.New("could not parse id from query")
	},
}
