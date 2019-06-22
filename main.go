package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)
// Command line flags
var port = flag.String("port", "8081", "port on which application will run")

const usage = `Diginex Micro Service 2.

Usage:

  dignex-micro-service-2 [commands|flags]

The commands & flags are:

  help              prints help

  --port              port on which application will run (default: 8081)

Examples:

  # prints help:
  dignex-micro-service-2 help

  # sample usage
  dignex-micro-service-2 --port 8081
`
// Show usage message and exit
func usageExit(rc int) {
	fmt.Println(usage)
	os.Exit(rc)
}

// Validate the command line flags
func validate() {
	if _, err := strconv.Atoi(*port); err != nil {
		log.Printf("ERROR: Please enter a valid port number.\n\n")
		usageExit(1)
	}
}

// Entrypoint
func main() {

	// Parse flags
	flag.Usage = func() { usageExit(0) }
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		switch args[0] {
		case "help":
			usageExit(0)
			return
		}
	}

	// Validate flags
	validate()

	// Register API endpoints
	http.HandleFunc("/reverse", handler)

	// Listen and serve API's
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Printf("ERROR: %q", err)
	}
}
