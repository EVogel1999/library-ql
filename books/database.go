package books

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"library-ql/database"
	"time"
)

type Book struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ISBN        string   `json:"isbn"`
	Download    string   `json:"download"`
	Authors     []string `json:"authors"`
}

var books *mongo.Collection

func configureDatabase() {
	books = database.Client.Database(database.Database).Collection(database.Books)
}

func getBookByID(id string) (*Book, error) {
	var result *Book

	// Perform search
	filter := bson.D{{"_id", id}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := books.FindOne(ctx, filter).Decode(&result)

	// Check if the author was found or if there was an error
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("book not found")
	} else if err != nil {
		return nil, err
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
	fmt.Print(authorID)
	cursor, err := books.Find(ctx, filter)
	fmt.Print("Wut")
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
