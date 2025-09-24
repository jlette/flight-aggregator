/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	controllers "aggregator/controllers"
	"fmt"
	"net/http"
)

func main() {


	controllers.Routes()
	fmt.Println("Starting server on :3001")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
    fmt.Println("Error starting server:", err)
	}
}
