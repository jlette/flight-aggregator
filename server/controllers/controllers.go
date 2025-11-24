package controllers

import (
	handlers "aggregator/http"
	"net/http"
)

// configure toutes les routes de notre API.
func Routes(){
	
	http.HandleFunc("/",handlers.Home) // Page d'accueil
	http.HandleFunc("/health",handlers.Healt) // Vérifie si le serveur tourne bien
	http.HandleFunc("/flight",handlers.FlightSorted) // Récupère tous les vols triés
}