package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	debug = flag.Bool("debug", false, "Show debug information")
)

func ini() {
	flag.Parse()
	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(new(log.TextFormatter))
}

func main() {
	log.Info("Starting atsregistry...")
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}
