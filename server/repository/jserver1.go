package repository

import (
	"aggregator/models"
	"encoding/json"
	"strings"
	"time"
)
type FlightsDTO struct {
	BookingId int `json:"bookingId"`
	Status string `json:"status"`
    PassengerName string `json:"passengerName"`
    FlightNumber int `json:"flightNumber"`
    DepartureAirport string `json:"departureAirport"`
    ArrivalAirport string `json:"arrivalAirport"`
    DepartureTime time.Time `json:"departureTime"`
    ArrivalTime time.Time `json:"arrivalTime"`
    Price int `json:"price"`
    Currency string `json:"currency"`
    Id int `json:"id"`

}

func (flightsDTO FlightsDTO) SetFlights() ([]models.Flight) {
	parts := strings.Fields(flightsDTO.PassengerName)
	first, last := parts[0], strings.Join(parts[1:], " ")
	flight := models.Flight{
		BookingId:       flightsDTO.BookingId,
		Status:          flightsDTO.Status,
		Passenger: models.Passenger{
				FirstName: first,
				LastName: last,
		},
		FlightNumber:    flightsDTO.FlightNumber,
		DepartureAirport: flightsDTO.DepartureAirport,
		ArrivalAirport:   flightsDTO.ArrivalAirport,
		DepartureTime:   flightsDTO.DepartureTime,
		ArrivalTime:     flightsDTO.ArrivalTime,
		Price:          flightsDTO.Price,
		Currency:       flightsDTO.Currency,
		Id:             flightsDTO.Id,
	}
	return []models.Flight{flight}

}

func GetFlights(data []byte) ([]FlightsDTO, error) {
	var flightsDTOs []FlightsDTO
	err := json.Unmarshal(data, &flightsDTOs)	
	if err != nil {
		return nil, err
	}
	return flightsDTOs, nil
}