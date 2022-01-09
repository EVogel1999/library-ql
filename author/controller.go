package author

import (
	"errors"
	"github.com/graphql-go/graphql"
	"library-ql/books"
)

func InitController() {
	configureDatabase()
}

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
			Type: graphql.NewList(books.BookType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				authorID := p.Source.(Author).ID
				if authorID != "" {
					books, err := books.GetBooksByAuthor(authorID)
					if err != nil {
						return nil, err
					}
					return books, nil
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
			author, err := getAuthorByID(id)
			if err != nil {
				return nil, err
			}
			return author, nil
		}
		return nil, errors.New("could not parse id from query")
	},
}
