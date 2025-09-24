/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	handlers "aggregator/http"
	"fmt"
	"io"
	"net/http"
)

func main() {

	home := func (w http.ResponseWriter,_ *http.Request ) {
		io.WriteString(w, "Home Page\n")
	}

	http.HandleFunc("/",home)
	http.HandleFunc("/health",handlers.Healt)
	http.HandleFunc("/flight",handlers.Flight)

	fmt.Println("Starting server on :3001")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
    fmt.Println("Error starting server:", err)
	}
}
