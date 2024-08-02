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
	"github.com/rs/cors"
)

// var db_client *mongo.Client

func main() {
	err := godotenv.Load()
	if err != nil {
		// panic(err)
	}
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.Handle("/api/", middleware.AuthorizationMiddleware(http.StripPrefix("/api", (routes.RoutesMain()))))
	libs.DBConnect()
	fmt.Println("Running on port ", port)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
