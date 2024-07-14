package libs

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db_client *mongo.Client

func DBConnect() {
	// Get the MongoDB URL from the environment variable
	db_url := os.Getenv("DATABASE_HOST")
	if db_url == "" {
		// Error handling if the DATABASE_HOST environment variable is not set
		fmt.Println("DATABASE_HOST environment variable not set")
		panic("DATABASE_HOST environment variable not set")
	}

	// Create server API options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Create MongoDB client options with the provided URL and server API options
	opts := options.Client().ApplyURI(db_url).SetServerAPIOptions(serverAPI)

	// Attempt to connect to the MongoDB server
	var err error
	db_client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		// Handle error if the connection fails
		fmt.Println("Database Connection Error !!!")
		panic(err)
	}
	// Successfully connected to MongoDB
	fmt.Println("Connected to MongoDB!")
}
func DBCollection(db string) *mongo.Collection {
	if db_client == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	collection := db_client.Database("umikyodb").Collection(db)
	return collection
}
