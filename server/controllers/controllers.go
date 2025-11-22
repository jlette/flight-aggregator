package controllers

import (
	handlers "aggregator/http"
	"net/http"
)

func Routes(){
	
	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/health",handlers.Healt)
	http.HandleFunc("/flight",handlers.Flight)
}