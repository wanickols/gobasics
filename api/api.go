package api

import (
	"encoding/json"
	"net/http"
)

// Coin balance parameters
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	//Success Code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	//Error code
	Code int

	//Error message
	Message string
}

// Return error response to person who called the endpoint
func writeError(w http.ResponseWriter, message string, code int) {

	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// Wrappers for the private function above
var (
	// User errors
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	//Server side errors
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)
