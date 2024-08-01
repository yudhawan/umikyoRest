package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	routes "umikyoRest/api"
	"umikyoRest/libs"
	"umikyoRest/middleware"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var db_client *mongo.Client

func main() {
	err := godotenv.Load()
	if err != nil {
		// log.Fatalf("Error .env file: %v", err)
	}
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.Handle("/api/", middleware.AuthorizationMiddleware(http.StripPrefix("/api", (routes.RoutesMain()))))
	libs.DBConnect()
	fmt.Println("Running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
