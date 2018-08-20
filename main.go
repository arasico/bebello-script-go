package main

import (
	"bebello-script-go/routes"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load(os.Getenv("GOPATH") + "/src/bebello-script-go/.env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := routes.NewRouter()
	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	fmt.Printf("Serving at localhost:%s...", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
