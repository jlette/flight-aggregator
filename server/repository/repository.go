package repository

import (
	"aggregator/models"
)
type FlightRepository interface {

	SetFlights() ([]models.Flight)
}	