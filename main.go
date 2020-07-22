package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server
	router := mux.NewRouter()

	router.HandleFunc("/singleplayer", insertSingleplayer).Methods("POST")
	router.HandleFunc("/multiplayer", insertMultiplayer).Methods("POST")
	router.HandleFunc("/singleplayer", getSingleplayerLeaderboard).Methods("GET")
	router.HandleFunc("/multiplayer", getMultiplayerLeaderboard).Methods("GET")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", router))
}
