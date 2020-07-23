package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// Hello world, the web server
	router := mux.NewRouter()
	router.Use(GenericMiddleware)

	router.HandleFunc("/singleplayer", InsertSingleplayer).Methods("POST")
	router.HandleFunc("/multiplayer", InsertMultiplayer).Methods("POST")
	router.HandleFunc("/singleplayer", GetSingleplayerLeaderboard).Methods("GET").Queries("sortBy", "{[a-zA-Z]+}")
	router.HandleFunc("/multiplayer", GetMultiplayerLeaderboard).Methods("GET").Queries("sortBy", "{[a-zA-Z]+}")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
