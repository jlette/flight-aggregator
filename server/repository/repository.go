package repository

import (
	"aggregator/models"
)
type FlightRepository interface {

	TransformFlights() ([]models.Flight)
}	