package repository

import (
	"aggregator/models"
)

// FlightRepository définit l’interface que doivent suivre nos deux repositories (j-server1 et j-server2).
type FlightRepository interface {
	// Renvoie les vols au format models.Flight, peu importe si ça vient de j-server1 ou j-server2.
	SetFlights() ([]models.Flight)
}	