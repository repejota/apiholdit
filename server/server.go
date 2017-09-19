// Copyright 2017 The apiholdit Authors. All rights reserved.

package server

import (
	"fmt"
	"net/http"

	"github.com/repejota/apiholdit/routes"
)

// Start starts the HTTP server for the qurl API microservice.
func Start(address string, port string) {

	http.HandleFunc("/teapot", routes.TeaPot)
	http.HandleFunc("/i", routes.PlaceHolder)

	// Start server
	serveraddress := fmt.Sprintf("%s:%s", address, port)
	http.ListenAndServe(serveraddress, nil)
}
