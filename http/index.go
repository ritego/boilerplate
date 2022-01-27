package http

import (
	"log"
	"net/http"
	"time"

	"github.com/aellacredit/jara/config"
)

func Init() {
	addr := config.String("SERVER_PORT")

	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: time.Second * time.Duration(config.Integer("SERVER_WRITE_TIMEOUT")),
		ReadTimeout:  time.Second * time.Duration(config.Integer("SERVER_READ_TIMEOUT")),
	}

	log.Printf("Server running on: %s", addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	routes()
}
