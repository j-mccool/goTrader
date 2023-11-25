package alpacaApi

import (
	"encoding/json"
	"goTrader/internal/config"
	"net/http"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
)

type AlpacaAPIImpl struct {
	client *alpaca.Client
}

func NewAplacaAPIImpl(c config.Config) AlpacaAPIImpl {
	return AlpacaAPIImpl{
		client: newClient(c),
	}
}

//Create new alpaca client
func newClient(c config.Config) *alpaca.Client{
	return alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    c.AlpacaAPIKey,
		APISecret: c.AlpacaAPISecret,
		BaseURL:   c.AlpacaBaseUrl,
	})
} 

func (i *AlpacaAPIImpl) GetAccountInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account, err := i.client.GetAccount()
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