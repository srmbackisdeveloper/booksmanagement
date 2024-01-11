package common

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var db *mongo.Database

func GetDbCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func InitDb() error {
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		return errors.New("You must set your 'MONGODB_URI' env variable. See\n\t https...")
	}

	client, err := mongo.Connect(context.Background(), options.Client(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	db = client.Database("go_demo")

	return nil
}

func CloseDb() error {
	return db.Client().Disconnect(context.Background())
}
