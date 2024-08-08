package libs

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Field struct {
	FieldName string
}

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

func Collection(table string) *mongo.Collection {
	collection := DBCollection(table)
	if collection == nil {
		fmt.Printf("Couldn't find collection of %v", table)
		// SendErrorResponse(w , "Database collection not found", http.StatusInternalServerError)
	}

	// cursor, err := collection.Find(context.Background(), user)
	// if err != nil {
	// 	SendErrorResponse(w, "Error Database Collection Users", 502)
	// }
	// defer cursor.Close(context.Background())
	return collection
}

var data []bson.M

func (c *Field) GetOne(key string, val any) *mongo.SingleResult {
	collection := Collection(c.FieldName)
	cursor := collection.FindOne(context.Background(), bson.D{{key, val}})
	return cursor
}
func (c *Field) GetMany(key string, val any) []bson.M {
	collection := Collection(c.FieldName)
	cursor, err := collection.Find(context.Background(), bson.D{{key, val}})
	if err != nil {
		fmt.Print(err)
	}
	cursor.All(context.Background(), &data)
	return data
}
func (c *Field) GetAll() []bson.M {
	collection := Collection(c.FieldName)
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Print(err)
	}
	cursor.All(context.Background(), &data)
	return data
}
func (c *Field) Insert(datum any) (*mongo.InsertOneResult, error) {
	collection := Collection(c.FieldName)
	cursor, err := collection.InsertOne(context.Background(), datum)
	if err != nil {
		fmt.Printf("Insert function : %v", err)
	}
	return cursor, err
}
