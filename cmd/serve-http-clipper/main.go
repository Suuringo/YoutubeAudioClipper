package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Suuringo/voice-clip-studio/server"
)

func main() {
	var port = 3030
	var allowedOrigins = "*"

	flagSet := flag.NewFlagSet("cmd", flag.ExitOnError)
	flagSet.IntVar(&port, "port", 3030, "set the port the server listens to")
	flagSet.StringVar(&allowedOrigins, "origins", allowedOrigins, "set the allowed origins to put in Access-Control-Allow-Origin header")

	flagSet.Parse(os.Args[1:])

	var stringPort = fmt.Sprintf("%v", port)

	log.Printf("Listening on port %v with origins '%v'\n", stringPort, allowedOrigins)
	mux := server.GetServer()
	server.AllowedOrigins = allowedOrigins
	log.Fatal(http.ListenAndServe(":"+stringPort, mux))
}
