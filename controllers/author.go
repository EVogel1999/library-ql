package controllers

import (
	"errors"
	"github.com/graphql-go/graphql"
	"library-ql/database"
)

var authorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
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
		"books": &graphql.Field{
			Type: graphql.NewList(bookType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorID := p.Source.(database.Author).ID
				if authorID != "" {
					if books, err := database.GetBooksByAuthor(authorID); err != nil {
						return nil, err
					} else {
						return books, nil
					}
				}
				return nil, errors.New("could not parse author id from query")
			},
		},
	},
})

var FindAuthorById = &graphql.Field{
	Type:        authorType,
	Description: "Gets a single author",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(string)
		if ok {
			author, err := database.GetAuthorByID(id)
			if err != nil {
				return nil, err
			}
			return author, nil
		}
		return nil, errors.New("could not parse id from query")
	},
}
