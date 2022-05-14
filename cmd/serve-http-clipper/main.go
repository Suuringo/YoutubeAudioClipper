package main

import (
	"log"
	"net/http"

	"github.com/Suuringo/voice-clip-studio/server"
)

func main() {
	log.Println("Listening on port 3030")

	mux := server.GetServer()

	log.Fatal(http.ListenAndServe(":3030", mux))
}
