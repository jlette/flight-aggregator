/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	handlers "aggregator/http"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/healt",handlers.Healt)

	http.ListenAndServe(":3000",nil)
	fmt.Println("starting point !")
}
