package main

import (
	"fmt"
	alpacaApi "goTrader/internal/api/alpaca"
	"goTrader/internal/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	alpApi := alpacaApi.NewAplacaAPIImpl(config)

	// Initialize Gin router
	router := mux.NewRouter()
	// router.HandleFunc("/start-stream", startStreamHandler).Methods("GET")
	router.HandleFunc("/accountInfo", alpApi.GetAccountInfoHandler()).Methods("GET")

	fmt.Printf("Starting server on port %s...", config.Port)
	log.Fatal(http.ListenAndServe(":5000", router))
}

// func startStreamHandler(w http.ResponseWriter, r *http.Request) {
// 	account, err := alpaca.GetAccount()
// 	if err != nil {
// 		log.Fatal("Error getting account:", err)
// 	}

// 	fmt.Printf("Connected to account: %s\n", account.ID)

// 	conn, err := alpaca.Stre
// }
