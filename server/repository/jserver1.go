package repository

import (
	"aggregator/models"
	"encoding/json"
	"strings"
	"time"
)

/* DTO utilisé par j-server1 pour représenter une réservation de vols.
On convertis ça ensuite en models.Flight grâce à la fonction SetFlights().*/
type FlightsDTO struct {
	BookingId string `json:"bookingId"`
	Status string `json:"status"`
    PassengerName string `json:"passengerName"`
    FlightNumber string `json:"flightNumber"`
    DepartureAirport string `json:"departureAirport"`
    ArrivalAirport string `json:"arrivalAirport"`
    DepartureTime time.Time `json:"departureTime"`
    ArrivalTime time.Time `json:"arrivalTime"`
    Price int `json:"price"`
    Currency string `json:"currency"`
    Id string `json:"id"`

}
/* adapte les données de j-server1 pour les rendre compatibles
avec notre modèle Flight commun au deux serveurs. */
func (f FlightsDTO) SetFlights() ([]models.Flight) {
	parts := strings.Fields(f.PassengerName)
	first, last := parts[0], strings.Join(parts[1:], " ")
	flight := models.Flight{
		BookingId:       f.BookingId,
		Status:          f.Status,
		Passenger: models.Passenger{
				FirstName: first,
				LastName: last,
		},
		FlightNumber:    f.FlightNumber,
		DepartureAirport: f.DepartureAirport,
		ArrivalAirport:   f.ArrivalAirport,
		DepartureTime:   f.DepartureTime,
		ArrivalTime:     f.ArrivalTime,
		Price:          f.Price,
		Currency:       f.Currency,
		Id:             f.Id,
	}
	return []models.Flight{flight}

}
//Convertit les données JSON venant de j-server1 en tableau de FlightsDTO pour réutiliser ensuite dans la fonction SetFlights()
func GetFlights(data []byte) ([]FlightsDTO, error) {
	var flightsDTOs []FlightsDTO
	err := json.Unmarshal(data, &flightsDTOs)	
	if err != nil {
		return nil, err
	}
	return flightsDTOs, nil
}