package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"umikyoRest/libs"
	"umikyoRest/routes"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var db_client *mongo.Client

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", routes.RoutesMain()))
	libs.DBConnect()
	fmt.Println("Running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
