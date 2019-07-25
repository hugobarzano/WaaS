package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"waas/controller"
	"fmt"
)

func main() {

	// Get the "PORT" env variable
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Print("$PORT must be set\n")
		port="8080"

	}

	fmt.Println("WaaS End-Points::")
	router := controller.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	fmt.Println("WaaS Listening on ::"+port)
	router.HandleFunc("/waas",controller.CommunityWebSocket)
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}