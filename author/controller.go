package author

import (
	"errors"
	"github.com/graphql-go/graphql"
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
