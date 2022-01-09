package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AuthorDatabase struct{}

type Author struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Pseudonym   string `json:"pseudonym"`
}

var authors *mongo.Collection

func (a AuthorDatabase) configureDatabase() {
	authors = client.Database(database_name).Collection(authors_col)
}

func GetAuthorByID(id string) (Author, error) {
	var result Author

	// Perform search
	filter := bson.D{{"_id", id}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := authors.FindOne(ctx, filter).Decode(&result)

	// Check if the author was found or if there was an error
	if err == mongo.ErrNoDocuments {
		return Author{}, errors.New("author not found")
	} else if err != nil {
		return Author{}, err
	}

	// Return result
	return result, nil
}

func GetAuthorsByIDs(authorIDs []string) ([]Author, error) {
	var result []Author

	// Perform search
	filter := bson.M{"_id": bson.M{"$in": authorIDs}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := authors.Find(ctx, filter)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("authors for given book not found")
	} else if err != nil {
		return nil, err
	}

	// Parse the cursor, if there is an error, send that
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	// Return result
	return result, nil
}
