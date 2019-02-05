package main

import (
	"fmt"
	"github.com/mstruebing/go-examples/11_webserver/logger"
	"log"
	"net/http"
	"os"
	"strconv"
)

const DEFAULT_PORT = 3000

// Returns either the DEFAULT_PORT or a port specified via environment variables
func getPort() string {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	// strconv.Atoi("") == 0
	if port == 0 || err != nil {
		port = DEFAULT_PORT
	}

	return fmt.Sprintf(":%d", port)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from our Webserver!\n")
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "This is a GET request")
	case "POST":
		fmt.Fprintf(w, "This is a OPTIONS request")
	case "OPTIONS":
		fmt.Fprintf(w, "This is a OPTIONS request")
	default:
		fmt.Fprintf(w, "This is an unsported request method")
		logger.Log(fmt.Sprintf("ERROR: Unsupported request made: %s", r.Method))
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/endpoint", endpointHandler)
	log.Fatal(http.ListenAndServe(getPort(), nil))
}
