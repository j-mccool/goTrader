package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// Config structure to hold the configuration data
type Config struct {
	AlpacaAPIKey    string `json:"apiKey"`
	AlpacaAPISecret string `json:"apiSecret"`
	AlpacaBaseUrl   string `json:"apiEndpoint"`
	Port            string `json:"port"`
}

func main() {
	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	//Create new alpaca client
	aClient := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    config.AlpacaAPIKey,
		APISecret: config.AlpacaAPISecret,
		BaseURL:   config.AlpacaBaseUrl,
	})

	// Initialize Gin router
	router := mux.NewRouter()
	// router.HandleFunc("/start-stream", startStreamHandler).Methods("GET")
	router.HandleFunc("/accountInfo", getAccountInfoHandler(aClient)).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}

// loadConfig loads the configuration from a JSON file
func loadConfig() (c Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	readerr := viper.ReadInConfig()

	if readerr != nil {
		fmt.Println("Error reading in config", readerr)
	}
	if c.Port = viper.GetString("Port"); err != nil {
		return
	}
	if c.AlpacaAPIKey = viper.GetString("apiKey"); err != nil {
		return
	}
	if c.AlpacaAPISecret = viper.GetString("apiSecret"); err != nil {
		return
	}
	if c.AlpacaBaseUrl = viper.GetString("apiEndpoint"); err != nil {
		return
	}
	return c, nil
}

// func startStreamHandler(w http.ResponseWriter, r *http.Request) {
// 	account, err := alpaca.GetAccount()
// 	if err != nil {
// 		log.Fatal("Error getting account:", err)
// 	}

// 	fmt.Printf("Connected to account: %s\n", account.ID)

// 	conn, err := alpaca.Stre
// }

func getAccountInfoHandler(aClient *alpaca.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account, err := aClient.GetAccount()
		if err != nil {
			http.Error(w, "Error getting account info", http.StatusInternalServerError)
			return
		}

		accountJSON, err := json.Marshal(account)
		if err != nil {
			http.Error(w, "error converting accounti nformation to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(accountJSON)
	}
}
