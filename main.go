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

type Config struct {
	Alpaca struct {
		APIKey      string `json:"apiKey"`
		APISecret   string `json:"apiSecret"`
		APIEndpoint string `json:"apiEndpoint"`
	} `json:"alpaca"`
}

func main() {
	config := loadConfig()
	// fmt.Println("It is a start...")

	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    config.Alpaca.APIKey,
		APISecret: config.Alpaca.APISecret,
		BaseURL:   config.Alpaca.APIEndpoint,
	})

	acct, err := client.GetAccount()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%_v\n", *acct)

	router := mux.NewRouter()

	router.HandleFunc("/account", getAccountInfo).Methods("GET")

	// Start server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}

func loadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %s", err)
	}

	return config

}

func getAccountInfo(w http.ResponseWriter, r *http.Request) {
	account, err := alpaca.GetAccount()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting account info %s", err), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(account)
	if err != nil {
		httpError(w, fmt.Sprintf("Error marshalling response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func httpError(w http.ResponseWriter, s string, i int) {
	panic("unimplemented")
}
