package main

import (
	"fmt"
	"net/http"

	"github/wanickols/gobasics/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)        //gets file and line number
	var r *chi.Mux = chi.NewRouter() //struct used to set up api
	handlers.Handler(r)              //Our handler function to set up router

	fmt.Println("Starting GO API Service...")

	fmt.Println(`
	___   ____   _____   _____   ___   ____
	/ _ \ |_  _| |_   _| | ____| |_ _| |  _ \
   | | | |  ||_    | |   |  _|    | |  | |_) |
   | |_| |   \ \   | |   | |___   | |  |  _ <
	\___/   _/ /   |_|   |_____| |___| |_| \_\

	`)

	err := http.ListenAndServe("localhost:8000", r) //start server

	if err != nil {
		log.Error(err) //logs any errors
	}

}
