package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CollectionDatabase interface {
	configureDatabase()
}

var client *mongo.Client

const (
	database_name = "library-ql"
	authors_col   = "authors"
	books_col     = "books"
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(errors.New("cannot connect to database"))
	}

	configureConnections()
}

func configureConnections() {
	AuthorDatabase{}.configureDatabase()
	BookDatabase{}.configureDatabase()
}
