package http

import (
	"log"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router = mux.NewRouter()
)

func routes() {

	router.HandleFunc("/", controller.home).Methods("GET")
	router.HandleFunc("/health", controller.health).Methods("GET")
	router.HandleFunc("/settlement", controller.settlement).Methods("POST")
	router.HandleFunc("/payout", controller.payout).Methods("POST")

	log.Println("Router Loaded")
}
