package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github/wanickols/gobasics/internal/middleware"
)

// Captial function is public

func Handler(r *chi.Mux) {
	//Global middleware (applied to all endpoints)
	r.Use(chimiddle.StripSlashes) //Prevents errors like https:localhost:8000/account/coins/ <-- last slash throwing 404

	//Account Route
	r.Route("/account", func(router chi.Router) {

		//Middleware for /account route to check for authorization
		router.Use((middleware.Authorization))

		router.Get("/coins", GetCoinBalance) //get returns specific function
	})
}
