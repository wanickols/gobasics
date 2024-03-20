package middleware

import (
	"errors"
	"github/wanickols/gobasics/api"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/wanickols/gobasics/api"
	"github.com/wanickols/gobasics/internals/tools"
)

var UnAuthorizedError = errors.New("Invalid username or token")

// middle ware is a function that gets called before the primary function
// This is a bouncer for the account route to check for authorization
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		//Get data from database

		//Connect to database
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		//Query database
		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*tokenDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return

		}

		next.ServeHTTP(w, r) //calls next middlerware or end function
	})
}
