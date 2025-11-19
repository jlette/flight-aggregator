package repository

import (
	"aggregator/models"
	"time"
)
type Flights struct {
	BookingId int `json: "bookingId"`
	Status string `json: "status"`
    PassengerName string `json: "passengerName"`
    FlightNumber string `json: "flightNumber"`
    DepartureAirport string `json: "departureAirport"`
    ArrivalAirport string `json: "arrivalAirport"`
    DepartureTime time.Time `json: "departureTime"`
    ArrivalTime time.Time `json: "arrivalTime"`
    price int `json: "price"`
    Currency string `json: "currency"`
    Id int `json: "id"`

}

func (f Flights) TransformFlights() ([]models.Flight) {
    
}