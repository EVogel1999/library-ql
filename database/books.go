package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BookDatabase struct{}

type Book struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ISBN        string   `json:"isbn"`
	Download    string   `json:"download"`
	Authors     []string `json:"authors"`
}

var books *mongo.Collection

func (b BookDatabase) configureDatabase() {
	books = client.Database(database_name).Collection(books_col)
}

func GetBookByID(id string) (Book, error) {
	var result Book

	// Perform search
	filter := bson.D{{"_id", id}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := books.FindOne(ctx, filter).Decode(&result)

	// Check if the author was found or if there was an error
	if err == mongo.ErrNoDocuments {
		return Book{}, errors.New("book not found")
	} else if err != nil {
		return Book{}, err
	}

	// Return result
	return result, nil
}

func GetBooksByAuthor(authorID string) ([]Book, error) {
	var result []Book

	// Perform search
	filter := bson.M{"authors": bson.M{"$in": [...]string{authorID}}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := books.Find(ctx, filter)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("books for given author not found")
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
