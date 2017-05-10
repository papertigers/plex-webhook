package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// CommandPath is the path to the command executed upon each plex webhook event
var (
	CommandPath    string
	CommandTimeout uint
)

func main() {
	// command line options
	listen := flag.String("listen", "127.0.0.1", "address to listen on")
	port := flag.String("port", "8080", "port to listen on")
	command := flag.String("command", "./event.sh", "path to the command that is execd upon each event")
	timeout := flag.Uint("timeout", 5, "amount of time in seconds to allow the command to run")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *listen, *port)
	router := NewRouter()
	CommandPath = *command
	CommandTimeout = *timeout
	log.Fatal(http.ListenAndServe(address, router))
}
