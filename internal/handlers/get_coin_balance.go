package handlers

import (
	"encoding/json"
	"github/wanickols/gobasics/api"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"github.com/wanickols/gobasics/api"
	"github.com/wanickols/gobasics/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var paramas = api.CoinBalanceParams{}
	var decoder *schema.decoder = schema.NewDecoder
	var err error

	err = decoder.Decode(&paramas, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(paramas.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
