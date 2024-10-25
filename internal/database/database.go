package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"time"
)

var (
	dburi = os.Getenv("DB_URI")
)

type Db *mongo.Client

func ConnectDB() *mongo.Client {
	// Opening a driver typically will not attempt to connect to the database.
	client, err := mongo.Connect(options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("MentorDB").Collection(collectionName)
	return collection
}
