// Copyright 2017 The apiholdit Authors. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/repejota/apiholdit"
	"github.com/repejota/apiholdit/server"
)

var (
	// Version is the latest release number of this microservice.
	//
	// This number is the latest tag from the git repository.
	Version string
	// Build is the lastest build string of this microservice.
	//
	// This string is the branch name and the commit hash (short format)
	Build string
)

func main() {
	// Default configuration
	ServerAddress := apiholdit.DefaultServerAddress
	ServerPort := apiholdit.DefaultServerPort

	// Read/Override configuration from environment variables
	EnvServerAddress := os.Getenv("API_SERVER_ADDRESS")
	if EnvServerAddress != "" {
		ServerAddress = EnvServerAddress
	}
	EnvServerPort := os.Getenv("API_SERVER_PORT")
	if EnvServerPort != "" {
		ServerPort = EnvServerPort
	}

	// Read/Override configuration from command line flags
	serverAddressPtr := flag.String("address", apiholdit.DefaultServerAddress, "Server address")
	serverPortPtr := flag.String("port", apiholdit.DefaultServerPort, "Server port")

	versionPtr := flag.Bool("version", false, "Show version information")
	flag.Parse()
	if *versionPtr {
		fmt.Println("apiholdit : Version", Version, "Build", Build)
		os.Exit(0)
	}
	if *serverAddressPtr != apiholdit.DefaultServerAddress {
		ServerAddress = *serverAddressPtr
	}
	if *serverPortPtr != apiholdit.DefaultServerPort {
		ServerPort = *serverPortPtr
	}

	address := fmt.Sprintf("%s:%s", ServerAddress, ServerPort)
	log.Println("Start qurl API server at", address, "...")
	server.Start(ServerAddress, ServerPort)
}
