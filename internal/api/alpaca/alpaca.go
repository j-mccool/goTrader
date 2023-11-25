package alpacaApi

import (
	"net/http"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
)

type AlpacaAPI interface {
	GetAccountInfoHandler(*alpaca.Client) (http.HandlerFunc)
}