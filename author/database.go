package author

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"library-ql/database"
	"time"
)

type Author struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Pseudonym   string `json:"pseudonym"`
}

var authors *mongo.Collection

func configureDatabase() {
	authors = database.Client.Database(database.Database).Collection(database.Authors)
}

func getAuthorByID(id string) (Author, error) {
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
